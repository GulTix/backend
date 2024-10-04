ALTER TABLE payments
ADD COLUMN issuer VARCHAR(255),
ADD COLUMN acquirer VARCHAR(255),
DROP COLUMN midtrans_payload;
