ALTER TABLE validations
DROP COLUMN payment_email_sent;

ALTER TABLE validations 
ADD COLUMN payment_email_sent_at TIMESTAMP;

