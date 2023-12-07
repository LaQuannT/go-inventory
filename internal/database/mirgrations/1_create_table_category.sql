-- +migrate Up
CREATE TABLE IF NOT EXISTS category (
    id SERIAL,
    name VARCHAR(50) PRIMARY KEY
  );

  -- +migrate Down
  DROP TABLE category;
