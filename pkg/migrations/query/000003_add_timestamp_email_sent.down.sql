ALTER TABLE validations 
DROP COLUMN payment_email_sent_at;

ALTER TABLE validations
ADD COLUMN payment_email_sent;