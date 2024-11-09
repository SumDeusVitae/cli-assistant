-- +goose Up
CREATE TABLE communications (
    id TEXT PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    model TEXT NOT NULL,
    question TEXT NOT NULL,
    reply TEXT,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE
    
);

-- +goose Down
DROP TABLE communications;