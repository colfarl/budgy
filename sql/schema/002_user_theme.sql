-- +goose Up
ALTER TABLE users
ADD theme TEXT DEFAULT "default";

-- +goose Down
ALTER TABLE users
DROP COLUMN theme;

