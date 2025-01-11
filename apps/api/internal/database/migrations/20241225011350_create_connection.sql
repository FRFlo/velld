-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE connections (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    host TEXT NOT NULL,
    port INTEGER NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    database_name TEXT NOT NULL,
    ssl INTEGER DEFAULT 0,
    database_size INTEGER,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP,
    last_connected_at TEXT,
    user_id TEXT REFERENCES users(id),
    status TEXT DEFAULT 'connected'
);

CREATE INDEX idx_connections_user_id ON connections(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE connections;
SELECT 'down SQL query';
-- +goose StatementEnd
