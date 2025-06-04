-- Создаём таблицу пользователей
CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    nickname VARCHAR(100) NOT NULL UNIQUE,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE
);

-- Создаём таблицу фильмов
CREATE TABLE IF NOT EXISTS films (
    film_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    genre VARCHAR(100),
    year INT CHECK (year > 1800 AND year <= EXTRACT(YEAR FROM CURRENT_DATE))
);

