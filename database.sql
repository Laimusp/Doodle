CREATE TABLE institutes (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    short_name VARCHAR(50)
);

CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    institute_id BIGINT REFERENCES institutes(id),
    full_name VARCHAR(255),
    age BIGINT,
    course BIGINT,
    email VARCHAR(255) UNIQUE NOT NULL,
    nickname VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);
