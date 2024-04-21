package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

func main() {
	// Строка подключения к базе данных (измените на свои данные)
	connStr := "postgres://project:project@localhost:5432/users?sslmode=disable"

	// Подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Создание таблицы Institutes
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Institutes (
			id BIGSERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			short_name VARCHAR(50) UNIQUE NOT NULL
		);
	`)
	if err != nil {
		panic(err)
	}

	// Создание таблицы Users
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Users (
			id BIGSERIAL PRIMARY KEY,
			institute_id BIGINT REFERENCES Institutes(id),
			full_name VARCHAR(255),
			age BIGINT,
			course BIGINT,
			email VARCHAR(255) UNIQUE NOT NULL,
			nickname VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL
		);
	`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Успешно подключено к PostgreSQL!")

	// Настройка маршрутов
	http.HandleFunc("/register", registerHandler(db)) // Передаем объект db в обработчик

	// Запуск веб-сервера
	fmt.Println("Сервер запущен на порту 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

// Обработчик регистрации
func registerHandler(db *sql.DB) http.HandlerFunc { // Принимаем db как аргумент
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "register.html")
		} else if r.Method == "POST" {
			nickname := r.FormValue("nickname")
			email := r.FormValue("email")
			password := r.FormValue("password")

			// Проверка, существует ли пользователь
			var exists bool

			err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE nickname = $1)", nickname).Scan(&exists)
			if err != nil {
				panic(err)
			}
			if exists {
				fmt.Fprintf(w, "Пользователь с именем %s уже существует", nickname)
				return
			}

			// Вставка нового пользователя (обратите внимание на безопасность паролей!)
			_, err = db.Exec("INSERT INTO users (nickname, email, password_hash) VALUES ($1, $2, $3)", nickname, email, password)
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, "Регистрация прошла успешно для пользователя %s", nickname)
		}
	}
}
