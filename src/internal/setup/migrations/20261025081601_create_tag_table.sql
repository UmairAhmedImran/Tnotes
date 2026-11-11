-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tags (
    tag_id TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    synced_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tags;
-- +goose StatementEnd