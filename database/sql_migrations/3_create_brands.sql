-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE brands (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(256) NOT NULL,
    created_at TIMESTAMP,
	updated_at TIMESTAMP
);

-- +migrate StatementEnd