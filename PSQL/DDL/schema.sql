DROP TABLE IF EXISTS characters;

CREATE TABLE characters (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  class varchar(255) NOT NULL,
  url varchar(255) NOT NULL,
  hp int NOT NULL,
  def int NOT NULL
);

\i /Users/RyH/goWork/src/github.com/hardr/rpgo-api/PSQL/DML/characters.sql;
