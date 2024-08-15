CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS forms (
    id VARCHAR(36) PRIMARY KEY,
    title TEXT NOT NULL
);

CREATE TYPE question_type_enum AS ENUM (
    'multiple_choice',
    'dropdown',
    'short_text',
    'long_text'
);

CREATE TABLE IF NOT EXISTS questions (
    id VARCHAR(36) PRIMARY KEY,
    type question_type_enum NOT NULL,
    question TEXT NOT NULL,
    form_id VARCHAR(36) NOT NULL,
    CONSTRAINT fk_questions_form_id FOREIGN KEY (form_id) REFERENCES forms(id)
);

CREATE TABLE IF NOT EXISTS answers (
    id VARCHAR(36) PRIMARY KEY,
    question_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    answer TEXT NOT NULL,
    CONSTRAINT fk_answers_question_id FOREIGN KEY (question_id) REFERENCES questions(id),
    CONSTRAINT fk_answers_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS options (
    id VARCHAR(36) PRIMARY KEY,
    question_id VARCHAR(36) NOT NULL,
    option TEXT NOT NULL,
    CONSTRAINT fk_options_question_id FOREIGN KEY (question_id) REFERENCES questions(id)
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_questions_form_id ON questions(form_id);
CREATE INDEX idx_answers_question_id ON answers(question_id);
CREATE INDEX idx_answers_user_id ON answers(user_id);
CREATE INDEX idx_options_question_id ON options(question_id);
