-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks (
    task_id TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    due_date DATE,
    status TEXT,
    priority TEXT,
    workspace_id INTEGER,
    is_private INTEGER DEFAULT 0
        CHECK (is_private IN (0,1)),
    is_deleted INTEGER DEFAULT 0
        CHECK (is_deleted IN (0, 1)),
    sort_order INTEGER DEFAULT 0,
    source TEXT,
    additional_info TEXT DEFAULT '{}',
    added_by_id TEXT NOT NULL,
    recurring_task_id TEXT,
    recurring_period TEXT,
    last_recurring_created_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,

    FOREIGN KEY (workspace_id) REFERENCES workspaces(workspace_id),
    FOREIGN KEY (added_by_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS tasks_versions (
    task_version_id TEXT PRIMARY KEY NOT NULL,
    task_id TEXT NOT NULL,
    task_snapshot TEXT NOT NULL,
    synced_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,

    FOREIGN KEY (task_id) REFERENCES tasks(task_id)
);

CREATE TABLE IF NOT EXISTS tasks_tags (
    task_tag_id TEXT PRIMARY KEY NOT NULL,
    task_id TEXT NOT NULL,
    tag_id TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,

    FOREIGN KEY (task_id) REFERENCES tasks(task_id),
    FOREIGN KEY (tag_id) REFERENCES tags(tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks_tags;
DROP TABLE IF EXISTS tasks_versions;
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd
