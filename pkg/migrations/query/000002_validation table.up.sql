-- CREATE TABLE IF NOT EXISTS validations (
--     id VARCHAR PRIMARY KEY,
--     answer_id VARCHAR,
--     user_id VARCHAR,
--     volunteer_id VARCHAR,
--     deleted BOOLEAN DEFAULT FALSE,
--     classification VARCHAR NOT NULL,
--     FOREIGN KEY (answer_id) REFERENCES answers(id),
--     FOREIGN KEY (user_id) REFERENCES users(id),
--     FOREIGN KEY (volunteer_id) REFERENCES volunteers(id)
-- );