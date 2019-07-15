BEGIN;
CREATE TABLE fund (
  id SERIAL, 
  fund_name VARCHAR (255), 
  price NUMERIC (8, 2),
  created TIMESTAMP
);
COMMIT;
SELECT version();