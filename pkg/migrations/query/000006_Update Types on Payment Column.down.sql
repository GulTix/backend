ALTER TABLE payments
DROP COLUMN acquirer,
ADD COLUMN update_at TIMESTAMP,
DROP COLUMN updated_at,
ADD COLUMN create_at TIMESTAMP,
DROP COLUMN created_at;
