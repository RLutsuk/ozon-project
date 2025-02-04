CREATE TABLE IF NOT EXISTS posts (
    id UUID        PRIMARY KEY DEFAULT uuid_generate_v4(),
    title          VARCHAR(255) NOT NULL,
    body           TEXT NOT NULL,
    allow_comments BOOLEAN DEFAULT TRUE,
    user_id UUID   REFERENCES users(id) ON DELETE CASCADE,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);          