-- +goose Up
-- SQL in this section is executed when the migration is applied.

ALTER TABLE equipments
  DROP COLUMN status;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

ALTER TABLE equipments
  ADD COLUMN status;