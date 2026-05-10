-- +goose Up
CREATE SCHEMA iam;

SET SEARCH_PATH TO iam, PUBLIC;

CREATE TABLE principals (
  id            text        NOT NULL,
  name          text        NOT NULL,
  email         text        NOT NULL,
  password      text        NOT NULL,
  avatar        text        NOT NULL,
  status        text        NOT NULL,
  roles         text        NOT NULL,
  customer_id   text        NOT NULL,
  provider      text        NOT NULL DEFAULT '',
  provider_user text        NOT NULL DEFAULT '',
  kind          text        NOT NULL,
  created_at    timestamptz NOT NULL DEFAULT NOW(),
  updated_at    timestamptz NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  UNIQUE (email)
);

CREATE TRIGGER created_at_principals_trgr
  BEFORE UPDATE
  ON principals
  FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
CREATE TRIGGER updated_at_principals_trgr
  BEFORE UPDATE
  ON principals
  FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

-- +goose Down
DROP SCHEMA IF EXISTS iam CASCADE;
