CREATE SCHEMA IF NOT EXISTS clicker;

CREATE TABLE IF NOT EXISTS clicker.users(
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS clicker.clicks(
    user_id INT NOT NULL REFERENCES clicker.users(id) ON DELETE CASCADE,
    click BIGINT
);