-- +migrate Up 
CREATE TABLE IF NOT EXISTS item (
  id SERIAL,
  sku TEXT UNIQUE PRIMARY KEY,
  name VARCHAR(255),
  brand VARCHAR(255),
  category VARCHAR(50) REFERENCES category(name),
  location VARCHAR(50));

-- migrate Down
DROP TABLE item;
