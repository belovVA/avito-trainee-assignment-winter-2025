CREATE TABLE purchases (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    merch_id INT NOT NULL,
    count INT NOT NULL DEFAULT 1 CHECK (count > 0),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (merch_id) REFERENCES merch(id) ON DELETE CASCADE,
    UNIQUE (user_id, merch_id) -- гарантирует, что у одного пользователя не будет дубликатов товаров
);
