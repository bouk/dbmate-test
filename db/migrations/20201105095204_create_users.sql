-- migrate:up

CREATE TABLE users (
  id BIGINT PRIMARY KEY,
  name TEXT NOT NULL
);


-- migrate:down

