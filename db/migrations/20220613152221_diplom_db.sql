-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(32) NOT NULL,
    second_name VARCHAR(32) NOT NULL,
    patronymic VARCHAR(32) NOT NULL,
    phone_number VARCHAR(32) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS admins(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    CONSTRAINT fk_id_user_id_admin
        FOREIGN KEY (user_id)
            REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS criminals(
    id SERIAL PRIMARY KEY,
    title_criminal VARCHAR(256),
    number_criminal VARCHAR(32)
);

CREATE TABLE IF NOT EXISTS offenses(
   id SERIAL PRIMARY KEY,
   user_id INT NOT NULL,
   crime_code_id INT NOT NULL,
   longitude VARCHAR(1024),
   latitude VARCHAR(1024),
   time timestamp default current_timestamp,
   description TEXT,
   CONSTRAINT fk_id_user_id
       FOREIGN KEY (user_id)
           REFERENCES users (id),
   CONSTRAINT fk_crime_code_id
       FOREIGN KEY (crime_code_id)
           REFERENCES criminals (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE offenses;

DROP TABLE admins;

DROP TABLE  criminals;

DROP TABLE users;
-- +goose StatementEnd
