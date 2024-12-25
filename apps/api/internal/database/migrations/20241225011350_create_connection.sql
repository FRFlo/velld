-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE connections (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    host VARCHAR(255) NOT NULL,
    port INTEGER NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    database_name VARCHAR(255) NOT NULL,
    ssl BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_connected_at TIMESTAMP,
    user_id UUID REFERENCES users(id),
    status VARCHAR(50) DEFAULT 'connected'
);

CREATE INDEX idx_connections_user_id ON connections(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE connections;
SELECT 'down SQL query';
-- +goose StatementEnd
