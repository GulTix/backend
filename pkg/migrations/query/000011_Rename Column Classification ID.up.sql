ALTER TABLE validations
DROP CONSTRAINT fk_classificationId,
DROP COLUMN classificationid;

ALTER TABLE validations
ADD COLUMN classification_id VARCHAR,
ADD CONSTRAINT fk_classification_id
FOREIGN KEY (classification_id)
REFERENCES classifications(id);
