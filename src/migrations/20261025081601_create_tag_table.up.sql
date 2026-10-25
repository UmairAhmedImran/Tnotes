-- Add up migration script here

CREATE TABLE IF NOT EXISTS tags (
    tag_id TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    synced_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

