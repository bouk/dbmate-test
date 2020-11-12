-- migrate:up

CREATE TABLE wuut (
  id BIGSERIAL PRIMARY KEY,
  uuid uuid NOT NULL DEFAULT uuid_generate_v4()
);


-- migrate:down

