-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE backups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    connection_id VARCHAR(255) REFERENCES connections(id),
    status VARCHAR(50) NOT NULL,
    path VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL,
    scheduled_time TIMESTAMP NOT NULL,
    completed_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE backups;
SELECT 'down SQL query';
-- +goose StatementEnd
