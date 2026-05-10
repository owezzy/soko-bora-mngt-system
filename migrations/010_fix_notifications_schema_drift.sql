-- +goose Up
CREATE SCHEMA IF NOT EXISTS notifications;

SET SEARCH_PATH TO notifications, public;

CREATE TABLE IF NOT EXISTS notifications (
  id          text        NOT NULL,
  order_id    text        NOT NULL,
  customer_id text        NOT NULL,
  type        text        NOT NULL,
  sms_number  text        NOT NULL,
  message     text        NOT NULL,
  created_at  timestamptz NOT NULL DEFAULT NOW(),
  updated_at  timestamptz NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS notifications_order_id_idx ON notifications (order_id);

DROP TRIGGER IF EXISTS created_at_notifications_trgr ON notifications;
CREATE TRIGGER created_at_notifications_trgr
  BEFORE UPDATE
  ON notifications
  FOR EACH ROW
EXECUTE PROCEDURE created_at_trigger();

DROP TRIGGER IF EXISTS updated_at_notifications_trgr ON notifications;
CREATE TRIGGER updated_at_notifications_trgr
  BEFORE UPDATE
  ON notifications
  FOR EACH ROW
EXECUTE PROCEDURE updated_at_trigger();

-- +goose Down
DROP TRIGGER IF EXISTS updated_at_notifications_trgr ON notifications.notifications;
DROP TRIGGER IF EXISTS created_at_notifications_trgr ON notifications.notifications;
DROP INDEX IF EXISTS notifications.notifications_order_id_idx;
DROP TABLE IF EXISTS notifications.notifications;
