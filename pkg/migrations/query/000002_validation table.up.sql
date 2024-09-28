CREATE TABLE IF NOT EXISTS validations (
    id VARCHAR PRIMARY KEY,
    answer_id VARCHAR,
    user_id VARCHAR,
    volunteer_id VARCHAR,
    deleted BOOLEAN DEFAULT FALSE,
    payment_email_sent BOOLEAN DEFAULT FALSE,
    classification VARCHAR NOT NULL,
    FOREIGN KEY (answer_id) REFERENCES answers(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (volunteer_id) REFERENCES volunteers(id)
);

-- Drop column bevy_link from events table
ALTER TABLE events
DROP COLUMN bevy_link;

-- Add column bevy_link to ticket_type table
ALTER TABLE ticket_type
ADD COLUMN bevy_link VARCHAR NOT NULL;  