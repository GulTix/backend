-- Remove the bevy_link column from ticket_type table
ALTER TABLE ticket_type
DROP COLUMN bevy_link;

-- Add bevy_link column back to events table
ALTER TABLE events
ADD COLUMN bevy_link VARCHAR NOT NULL  -- Adjust the data type as needed

-- Drop the validations table
DROP TABLE IF EXISTS validations;