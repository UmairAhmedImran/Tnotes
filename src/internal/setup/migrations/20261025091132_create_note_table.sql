-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS notes (
    note_id TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    workspace_id INTEGER,
    is_private INTEGER DEFAULT 0
        CHECK (is_private IN (0,1)),
    is_deleted INTEGER DEFAULT 0
        CHECK (is_deleted IN (0, 1)),
    sort_order INTEGER DEFAULT 0,
    source TEXT,
    additional_info TEXT DEFAULT '{}',
    added_by_id TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,

    FOREIGN KEY (workspace_id) REFERENCES workspaces(workspace_id),
    FOREIGN KEY (added_by_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS notes_versions (
    note_version_id TEXT PRIMARY KEY NOT NULL,
    note_id TEXT NOT NULL,
    note_snapshot TEXT NOT NULL,
    synced_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,

    FOREIGN KEY (note_id) REFERENCES notes(note_id)
);

CREATE TABLE IF NOT EXISTS notes_tags (
    note_tag_id TEXT PRIMARY KEY NOT NULL,
    note_id TEXT NOT NULL,
    tag_id TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,

    FOREIGN KEY (note_id) REFERENCES notes(note_id),
    FOREIGN KEY (tag_id) REFERENCES tags(tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notes_tags;
DROP TABLE IF EXISTS notes_versions;
DROP TABLE IF EXISTS notes;
-- +goose StatementEnd