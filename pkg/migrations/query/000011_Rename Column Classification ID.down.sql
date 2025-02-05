-- Down migration

-- Step 1: Remove the new foreign key constraint and column
ALTER TABLE validations
DROP CONSTRAINT IF EXISTS fk_classification_id,
DROP COLUMN IF EXISTS classification_id;

-- Step 2: Recreate the original foreign key constraint and column
ALTER TABLE validations
ADD COLUMN classificationid INTEGER,
ADD CONSTRAINT fk_classificationId
FOREIGN KEY (classificationid)
REFERENCES classifications(id);
```
