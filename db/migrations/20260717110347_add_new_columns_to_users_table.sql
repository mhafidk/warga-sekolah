-- +goose Up
ALTER TABLE users ADD COLUMN school_id INTEGER;
ALTER TABLE users ADD COLUMN role TEXT NOT NULL DEFAULT '';
ALTER TABLE users ADD COLUMN deleted_at DATETIME;

-- +goose Down
ALTER TABLE users DROP COLUMN school_id;
ALTER TABLE users DROP COLUMN role;
ALTER TABLE users DROP COLUMN deleted_at;
