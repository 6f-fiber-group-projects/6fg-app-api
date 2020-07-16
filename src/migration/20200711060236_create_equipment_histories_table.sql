-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS equipment_histories (
  id serial primary key,
  equip_id integer,
  user_id integer,
  reservation_id integer,
  start_date timestamp,
  end_date timestamp,
  created_at timestamp,
  updated_at timestamp
);

ALTER TABLE equipment_histories 
  ADD FOREIGN KEY (equip_id)
      REFERENCES equipments (id)
      ON DELETE SET NULL;

ALTER TABLE equipment_histories 
  ADD FOREIGN KEY (user_id)
      REFERENCES users (id)
      ON DELETE SET NULL;

ALTER TABLE equipment_histories 
  ADD FOREIGN KEY (reservation_id)
      REFERENCES equipment_reservations (id)
      ON DELETE SET NULL;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS equipment_histories;