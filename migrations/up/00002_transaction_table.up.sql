CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    from_user INT NOT NULL,
    to_user INT NOT NULL,
    amount INT NOT NULL DEFAULT 0 CHECK (amount >= 0),
    FOREIGN KEY (from_user) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (to_user) REFERENCES users(id) ON DELETE CASCADE
);
