-- +goose Up
-- +goose StatementBegin
CREATE TABLE notifications (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    title TEXT NOT NULL,
    message TEXT NOT NULL,
    type TEXT NOT NULL, -- 'backup_failed', 'backup_completed', etc.
    status TEXT NOT NULL DEFAULT 'unread', -- 'read', 'unread'
    metadata JSONB, -- Store additional data like backup_id, connection_id, etc.
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE notifications;
-- +goose StatementEnd
