CREATE TABLE IF NOT EXISTS users (
    id VARCHAR PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS events (
    id VARCHAR PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    bevy_link VARCHAR NOT NULL,
    google_form_link VARCHAR,
    deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS answers (
    id VARCHAR PRIMARY KEY,
    event_id VARCHAR,
    user_id VARCHAR,
    raw_data JSONB,
    deleted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (event_id) REFERENCES events(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS ticket_type (
    id VARCHAR PRIMARY KEY,
    "name" VARCHAR,
    price NUMERIC(19, 4),
    event_id VARCHAR,
    deleted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (event_id) REFERENCES events(id)
);

CREATE TABLE IF NOT EXISTS payments (
    id VARCHAR PRIMARY KEY,
    answer_id VARCHAR,
    ticket_type_id VARCHAR,
    qris_url VARCHAR,
    "status" VARCHAR,
    deleted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (answer_id) REFERENCES answers(id),
    FOREIGN KEY (ticket_type_id) REFERENCES ticket_type(id)
);

CREATE TABLE IF NOT EXISTS volunteers (
    id VARCHAR PRIMARY KEY,
    username VARCHAR NOT NULL UNIQUE,
    email VARCHAR NOT NULL UNIQUE,
    "role" VARCHAR,
    deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS volunteers_events (
    id VARCHAR PRIMARY KEY,
    volunteer_id VARCHAR,
    event_id VARCHAR,
    deleted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (volunteer_id) REFERENCES volunteers(id),
    FOREIGN KEY (event_id) REFERENCES events(id)
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_volunteer_username ON volunteers(username);
CREATE INDEX idx_payments_status ON payments("status");
