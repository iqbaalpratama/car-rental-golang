-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE customers (
    id SERIAL PRIMARY KEY, 
    first_name VARCHAR(256) NOT NULL,
    last_name VARCHAR(256) NOT NULL,
    email VARCHAR(256) UNIQUE NOT NULL,
    password VARCHAR(256) NOT NULL,
    address VARCHAR(256) NOT NULL,
    phone_number VARCHAR(256) NOT NULL,
    created_at TIMESTAMP,
	updated_at TIMESTAMP
);

-- +migrate StatementEnd