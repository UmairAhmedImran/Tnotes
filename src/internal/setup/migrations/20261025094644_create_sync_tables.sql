-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sync_history (
    sync_history_id TEXT PRIMARY KEY NOT NULL,
    "start" DATETIME NOT NULL,
    "end" DATETIME NOT NULL,
    metadata TEXT DEFAULT '{}',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS logs (
    log_id TEXT PRIMARY KEY NOT NULL,
    action TEXT NOT NULL,
    metadata TEXT DEFAULT '{}',
    status TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS providers (
    provider_id TEXT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    url TEXT,
    configuration TEXT DEFAULT '{}',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sync_history;
DROP TABLE IF EXISTS logs;
DROP TABLE IF EXISTS providers;
-- +goose StatementEnd