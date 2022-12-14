-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE cars (
    id SERIAL PRIMARY KEY, 
    car_number VARCHAR(256) NOT NULL,
    car_model VARCHAR(256) NOT NULL,
    car_year INT NOT NULL,
    rent_price INT NOT NULL,
    brand_id INT references brands(id),
    created_at TIMESTAMP,
	updated_at TIMESTAMP
);

-- +migrate StatementEnd