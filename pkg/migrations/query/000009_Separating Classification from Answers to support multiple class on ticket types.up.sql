CREATE TABLE classifications (
    id VARCHAR PRIMARY KEY,
    event_id VARCHAR,
    name VARCHAR NOT NULL,
    FOREIGN KEY (event_id) REFERENCES events(id)
);

CREATE TABLE allowed_classifications (
    id VARCHAR PRIMARY KEY,
    ticket_id VARCHAR,
    classification_id VARCHAR,
    FOREIGN KEY (ticket_id) REFERENCES ticket_type(id),
    FOREIGN KEY (classification_id) REFERENCES classifications(id)
);

ALTER TABLE validations
DROP COLUMN classification,
ADD COLUMN classificationId VARCHAR;

ALTER TABLE validations
ADD CONSTRAINT fk_classificationId
FOREIGN KEY (classificationId)
REFERENCES classifications(id);
