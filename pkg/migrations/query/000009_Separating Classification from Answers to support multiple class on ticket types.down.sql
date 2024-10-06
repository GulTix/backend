-- Remove the foreign key constraint from validations table
ALTER TABLE validations
DROP CONSTRAINT fk_classificationId;

-- Remove the classificationId column and add back the classification column
ALTER TABLE validations
DROP COLUMN classificationId,
ADD COLUMN classification VARCHAR;

-- Drop the allowed_classifications table
DROP TABLE IF EXISTS allowed_classifications;

-- Drop the classifications table
DROP TABLE IF EXISTS classifications;
