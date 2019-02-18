-- +migrate Up
CREATE TABLE somethings (id SERIAL PRIMARY KEY);

-- +migrate Down
DROP TABLE somethings;