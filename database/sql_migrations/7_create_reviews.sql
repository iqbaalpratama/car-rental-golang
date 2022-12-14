-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY, 
    rating INT NOT NULL,
    review VARCHAR(256) NOT NULL,
    transaction_id INT references transactions(id),
    created_at TIMESTAMP,
	updated_at TIMESTAMP
);

-- +migrate StatementEnd