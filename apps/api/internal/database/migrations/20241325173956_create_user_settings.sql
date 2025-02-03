-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_settings (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL UNIQUE,
    notify_dashboard BOOLEAN DEFAULT TRUE,
    notify_email BOOLEAN DEFAULT FALSE,
    notify_webhook BOOLEAN DEFAULT FALSE,
    webhook_url TEXT,
    email TEXT,
    smtp_host TEXT,
    smtp_port TEXT,
    smtp_username TEXT,
    smtp_password TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_settings;
-- +goose StatementEnd
