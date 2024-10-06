ALTER TABLE ticket_type
ADD COLUMN payment_expired_duration INT NOT NULL DEFAULT 24;
