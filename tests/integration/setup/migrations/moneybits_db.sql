CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    planner_id INTEGER NOT NULL,
    amount BIGINT NOT NULL,
    description TEXT NOT NULL,
    notes TEXT,
    attachments JSONB, 
    category VARCHAR(255) NOT NULL,
    month_year_index VARCHAR(7) NOT NULL, -- Format: YYYY-MM
    type VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    email TEXT 
);