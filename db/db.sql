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
                salary numeric(18, 2),
                payment_frequency varchar(10),
                monthly_tax numeric(18, 2),
                yearly_tax numeric(18, 2),
                "timestamp" timestamp,
                PRIMARY KEY (calculation_id),
                FOREIGN KEY (user_id) REFERENCES users (user_id)
);

INSERT INTO users(email, password) VALUES ('corne@gmail.com', 'asdf');
INSERT INTO users(email, password) VALUES ('asdf', 'asdf');
INSERT INTO calculations (calculation_id, user_id, salary, payment_frequency, monthly_tax, yearly_tax, timestamp) VALUES (1, 1, 20000.00, 'Monthly', 3693.33, 44319.96, current_timestamp(2));
INSERT INTO calculations (calculation_id, user_id, salary, payment_frequency, monthly_tax, yearly_tax, timestamp) VALUES (2, 2, 1400000.00, 'Yearly', 3321.08, 39852.96, CURRENT_TIMESTAMP(2));
INSERT INTO calculations (calculation_id, user_id, salary, payment_frequency, monthly_tax, yearly_tax, timestamp) VALUES (3, 2, 60000.00, 'Monthly', 16782.33, 201387.96, CURRENT_TIMESTAMP(2));
INSERT INTO calculations (calculation_id, user_id, salary, payment_frequency, monthly_tax, yearly_tax, timestamp) VALUES (4, 2, 2500000.00, 'Yearly', 6671.47, 80057.64, CURRENT_TIMESTAMP(2));
INSERT INTO calculations (calculation_id, user_id, salary, payment_frequency, monthly_tax, yearly_tax, timestamp) VALUES (5, 2, 1400000.00, 'Yearly', 3321.08, 39852.96, CURRENT_TIMESTAMP(2));