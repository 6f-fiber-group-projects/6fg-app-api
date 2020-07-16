-- +goose Up
-- SQL in this section is executed when the migration is applied.

ALTER TABLE users
  ADD access_token text;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

ALTER TABLE users
  DROP COLUMN access_token;