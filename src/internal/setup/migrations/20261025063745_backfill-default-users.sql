-- +goose Up
-- +goose StatementBegin
INSERT INTO users (user_id, username) VALUES
('00000000-0000-0000-0000-000000000001', 'system'),
('00000000-0000-0000-0000-000000000002', 'cron'),
('00000000-0000-0000-0000-000000000003', 'user');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users
WHERE user_id IN (
    '00000000-0000-0000-0000-000000000001',
    '00000000-0000-0000-0000-000000000002',
    '00000000-0000-0000-0000-000000000003'
);
-- +goose StatementEnd
