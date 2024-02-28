CREATE TABLE IF NOT EXISTS item (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  brand VARCHAR(255) NOT NULL,
  stock_keeping_unit TEXT UNIQUE NOT NULL,
  category VARCHAR(50),
  location VARCHAR(255),
  amount INT DEFAULT 0
);
