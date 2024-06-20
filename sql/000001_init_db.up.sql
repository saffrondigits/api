CREATE TABLE IF NOT EXISTS user_account (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(32),
    last_name VARCHAR(32),
    email VARCHAR(32) NOT NULL,
    password VARCHAR(32) NOT NULL
);
