CREATE TABLE IF NOT EXISTS users (
    id UUID    PRIMARY KEY DEFAULT uuid_generate_v4(),
    username   VARCHAR(50) NOT NULL UNIQUE,
    email      VARCHAR(100) NOT NULL UNIQUE,
    first_name VARCHAR(50),
    last_name  VARCHAR(50),
    created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP
);