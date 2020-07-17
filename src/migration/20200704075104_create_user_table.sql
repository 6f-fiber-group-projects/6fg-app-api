-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS users (
  id serial primary key,
  authority_id integer not null,
  google_id text,
  name varchar(255) not null unique,
  email varchar(255) not null unique,
  password text,
  created_at timestamp,
  updated_at timestamp
);

ALTER TABLE users
  ADD FOREIGN KEY (authority_id)
      REFERENCES authorities (id)
      ON DELETE SET NULL;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users;