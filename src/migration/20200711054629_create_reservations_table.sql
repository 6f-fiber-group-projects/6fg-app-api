-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS equipment_reservations (
  id serial primary key,
  equip_id integer not null,
  user_id integer not null,
  start_date timestamp,
  end_date timestamp,
  created_at timestamp,
  updated_at timestamp
);

ALTER TABLE equipment_reservations
  ADD FOREIGN KEY (equip_id)
      REFERENCES equipments (id)
      ON DELETE CASCADE;

ALTER TABLE equipment_reservations
  ADD FOREIGN KEY (user_id)
      REFERENCES users (id)
      ON DELETE CASCADE;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS equipment_reservations;