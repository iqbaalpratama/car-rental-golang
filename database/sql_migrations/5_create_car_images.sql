-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE car_images (
    id SERIAL PRIMARY KEY, 
    image_url VARCHAR(256) NOT NULL,
    car_id INT references cars(id),
    created_at TIMESTAMP,
	updated_at TIMESTAMP
);

-- +migrate StatementEnd