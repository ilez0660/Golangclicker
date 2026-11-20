CREATE SCHEMA IF NOT EXISTS clicker;

CREATE TABLE IF NOT EXISTS clicker.users(
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
    CHECK (
        email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'
    )
);
CREATE TABLE IF NOT EXISTS clicker.clicks(
    user_id INT NOT NULL REFERENCES clicker.users(id) ON DELETE CASCADE,
    click BIGINT
);