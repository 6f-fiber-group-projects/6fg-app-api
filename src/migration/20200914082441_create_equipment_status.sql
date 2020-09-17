-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS equipments_status (
  equip_id integer references equipments(id) ON DELETE CASCADE,
  user_id integer,
  status integer,
  created_at timestamp,
  updated_at timestamp
);

INSERT INTO equipments_status (equip_id, user_id, status) 
SELECT
    id,
    0, -- no user
    0 -- default status
FROM
    equipments
ORDER BY
    id;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS equipments_status;