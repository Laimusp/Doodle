CREATE TABLE "Institutes"(
    "id" bigserial NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "short_name" VARCHAR(50) NOT NULL
);
ALTER TABLE
    "Institutes" ADD PRIMARY KEY("id");
CREATE TABLE "Users"(
    "id" BIGINT NOT NULL,
    "institute_id" BIGINT NOT NULL,
    "full_name" VARCHAR(255) NOT NULL,
    "age" BIGINT NOT NULL,
    "course" BIGINT NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "nickname" VARCHAR(255) NOT NULL,
    "password_hash" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "Users" ADD PRIMARY KEY("id");
ALTER TABLE
    "Users" ADD CONSTRAINT "users_email_unique" UNIQUE("email");
ALTER TABLE
    "Users" ADD CONSTRAINT "users_nickname_unique" UNIQUE("nickname");
ALTER TABLE
    "Users" ADD CONSTRAINT "users_institute_id_foreign" FOREIGN KEY("institute_id") REFERENCES "Institutes"("id");