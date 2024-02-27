CREATE TABLE IF NOT EXISTS item (
  id SERIAL PRIMARY KEY,
  stock_keeping_unit TEXT NOT NULL,
  name VARCHAR(255) NOT NULL,
  brand VARCHAR(255) NOT NULL,
  category VARCHAR(50),
  location VARCHAR(50)
);
