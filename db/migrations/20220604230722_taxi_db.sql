-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL PRIMARY KEY,
    login VARCHAR(48) UNIQUE,
    password_hash VARCHAR(255)
    );

CREATE TABLE IF NOT EXISTS statuses(
    status_id SERIAL PRIMARY KEY,
    name VARCHAR(32)
    );

CREATE TABLE IF NOT EXISTS cab_mans(
    cab_man_id SERIAL PRIMARY KEY,
    first_name VARCHAR(32),
    second_name VARCHAR(32),
    vehicle_number VARCHAR(16) UNIQUE,
    image VARCHAR(255)
    );

CREATE TABLE IF NOT EXISTS orders(
    id SERIAL PRIMARY KEY,
    id_cab_man INT,
    id_status INT,
    start_location VARCHAR(64),
    end_location VARCHAR(64),
    created_at timestamp,
    updated_at timestamp default current_timestamp,
    CONSTRAINT fk_id_cab_man
        FOREIGN KEY (id_cab_man)
            REFERENCES cab_mans (cab_man_id),
    CONSTRAINT fk_status_order
        FOREIGN KEY (id_status)
            REFERENCES statuses (status_id)
);

CREATE TABLE IF NOT EXISTS orders_history(
    id SERIAL PRIMARY KEY,
    id_order INT,
    latitude VARCHAR(32),
    longitude VARCHAR(32),
    created_at timestamp default current_timestamp,
    CONSTRAINT fk_status_order
        FOREIGN KEY (id_order)
            REFERENCES orders (id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;

DROP TABLE orders_history;

DROP TABLE orders;

DROP TABLE statuses;

DROP TABLE cab_mans;

-- +goose StatementEnd
