DROP
    TABLE IF EXISTS calculations;
DROP
    TABLE IF EXISTS users;

CREATE TABLE users (
                user_id serial,
                email varchar(320) NOT NULL,
                password varchar(255) NOT NULL,
                PRIMARY KEY (user_id)
);

CREATE TABLE calculations (
                calculation_id serial,
                user_id int,
                payment_frequency varchar(10),
                salary float,
                tax float,
                "timestamp" timestamp,
                PRIMARY KEY (calculation_id),
                FOREIGN KEY (user_id) REFERENCES users (user_id)
);

INSERT INTO users(email, password) VALUES ('corne@gmail.com', 'asdf');
INSERT INTO users(email, password) VALUES ('asdf', 'asdf');