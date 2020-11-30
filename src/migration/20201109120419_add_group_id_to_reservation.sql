-- +goose Up
-- SQL in this section is executed when the migration is applied.

ALTER TABLE equipment_reservations
  ADD group_id text;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

ALTER TABLE equipment_reservations
  DROP COLUMN group_id;