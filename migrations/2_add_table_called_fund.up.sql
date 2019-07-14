BEGIN;
CREATE TABLE fund (
  id INTEGER, 
  fund_name VARCHAR (255), 
  price NUMERIC (8, 2),
  created TIMESTAMP
);
COMMIT;
SELECT version();