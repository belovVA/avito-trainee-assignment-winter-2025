CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    balance INT NOT NULL DEFAULT 1000
);
