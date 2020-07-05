-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS authorities (
  id serial primary key,
  name varchar(255) not null unique,
  created_at timestamp,
  updated_at timestamp
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS authorities;