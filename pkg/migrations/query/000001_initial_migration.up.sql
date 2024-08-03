CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    deleted BOOLEAN DEFAULT FALSE,
)

CREATE TABLE IF NOT EXIST forms (
    id VARCHAR(36) PRIMARY KEY,
    title TEXT NOT NULL,
)

CREATE ENUM question_type (
    'multiple_choice',
    'dropdown',
    'short_text',
    'long_text',
)

CREATE TABLE IF NOT EXISTS questions (
    id VARCHAR(36) PRIMARY KEY,
    type question_type NOT NULL,
    question TEXT NOT NULL,
    form_id VARCHAR(36) NOT NULL,
)

CREATE TABLE IF NOT EXISTS answers (
    id VARCHAR(36) PRIMARY KEY,
    question_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    answer TEXT NOT NULL,
    FOREIGN KEY (question_id) REFERENCES questions(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
    FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)

CREATE TABLE IF NOT EXISTS Options (
    id VARCHAR(36) PRIMARY KEY,
    question_id VARCHAR(36) NOT NULL,
    option TEXT NOT NULL,
    FOREIGN KEY (question_id) REFERENCES questions(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)
