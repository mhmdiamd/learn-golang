CREATE TABlE customers (
  id VARCHAR(100) primary key,
  name VARCHAR(100)
);

ALTER TABLE customers 
  ADD COLUMN email VARCHAR(100),
  ADD COLUMN balance int DEFAULT(0),
  ADD COLUMN rating numeric DEFAULT(0.0),
  ADD COLUMN created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  ADD COLUMN birth_date DATE,
  ADD COLUMN married BOOLEAN DEFAULT(false);

CREATE TABLE admin (
  id VARCHAR(100) primary key,
  username VARCHAR(100),
  password VARCHAR(100)
);

CREATE TABLE comments (
  id SERIAL primary key,
  email VARCHAR(100),
  comment text
);