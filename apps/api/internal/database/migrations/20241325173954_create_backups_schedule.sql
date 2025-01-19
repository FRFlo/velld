-- +goose Up
-- +goose StatementBegin
-- Create backup_schedules table
CREATE TABLE backup_schedules (
    id TEXT PRIMARY KEY,
    connection_id TEXT REFERENCES connections(id),
    enabled BOOLEAN DEFAULT FALSE,
    cron_schedule TEXT,
    retention_days INTEGER,
    next_run_time TEXT,
    last_backup_time TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE backups_schedules;
SELECT 'down SQL query';
-- +goose StatementEnd 