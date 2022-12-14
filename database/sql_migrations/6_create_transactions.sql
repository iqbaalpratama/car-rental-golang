-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY, 
    date_start DATE NOT NULL,
    date_finish DATE NOT NULL,
    total_price INT NOT NULL,
    status VARCHAR(256) NOT NULL,
    rating INT,
    review VARCHAR(256),
    car_id INT references cars(id),
    customer_id INT references customers(id),
    created_at TIMESTAMP,
	updated_at TIMESTAMP
);

-- +migrate StatementEnd