ALTER TABLE payments
ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
DROP COLUMN create_at,
ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
DROP COLUMN update_at,
ADD COLUMN acquirer VARCHAR;
