ALTER TABLE payments
DROP COLUMN issuer,
DROP COLUMN acquirer,
ADD COLUMN midtrans_payload JSONB;
