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
                salary numeric(18, 4),
                tax numeric(18, 4),
                "timestamp" timestamp,
                PRIMARY KEY (calculation_id),
                FOREIGN KEY (user_id) REFERENCES users (user_id)
);

INSERT INTO users(email, password) VALUES ('corne@gmail.com', 'asdf');
INSERT INTO users(email, password) VALUES ('asdf', 'asdf');
INSERT INTO calculations (calculation_id, user_id, payment_frequency, salary, tax, timestamp) VALUES (1, 1, 'Monthly', 20000.00, 3693.33, current_timestamp(2));
INSERT INTO calculations (calculation_id, user_id, payment_frequency, salary, tax, timestamp) VALUES (2, 2, 'Yearly', 1400000.00, 39853.00, CURRENT_TIMESTAMP(2));
INSERT INTO calculations (calculation_id, user_id, payment_frequency, salary, tax, timestamp) VALUES (3, 2, 'Monthly', 60000.00, 16782.33, CURRENT_TIMESTAMP(2));
INSERT INTO calculations (calculation_id, user_id, payment_frequency, salary, tax, timestamp) VALUES (4, 2, 'Yearly', 2500000.00, 80057.67, CURRENT_TIMESTAMP(2));
INSERT INTO calculations (calculation_id, user_id, payment_frequency, salary, tax, timestamp) VALUES (5, 2, 'Yearly', 1400000.00, 39853.00, CURRENT_TIMESTAMP(2));