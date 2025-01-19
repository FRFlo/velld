-- +goose Up
-- +goose StatementBegin
CREATE TABLE backups (
    id TEXT PRIMARY KEY,
    connection_id TEXT REFERENCES connections(id),
    schedule_id TEXT REFERENCES backup_schedules(id),
    status TEXT NOT NULL,
    path TEXT,
    size INTEGER,
    started_time TEXT,
    completed_time TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- Create indexes for better query performance
CREATE INDEX idx_backups_connection_id ON backups(connection_id);
CREATE INDEX idx_backups_schedule_id ON backups(schedule_id);
CREATE INDEX idx_backup_schedules_connection_id ON backup_schedules(connection_id);

-- +goose Down
-- +goose StatementBegin
DROP TABLE backups;
SELECT 'down SQL query';
-- +goose StatementEnd 