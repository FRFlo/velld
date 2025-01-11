-- +goose Up
-- +goose StatementBegin
CREATE TABLE backups (
    id TEXT PRIMARY KEY,
    connection_id TEXT REFERENCES connections(id),
    status TEXT NOT NULL,
    path TEXT NOT NULL,
    size INTEGER NOT NULL,
    scheduled_time TEXT NOT NULL,
    completed_time TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE backups;
SELECT 'down SQL query';
-- +goose StatementEnd
