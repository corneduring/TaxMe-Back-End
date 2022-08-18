DROP
    TABLE IF EXISTS calculations;
DROP
    TABLE IF EXISTS users;

CREATE TABLE users (
                user_id int,
                email varchar(320) NOT NULL,
                password varchar(255) NOT NULL,
                followers int,
                following int,
                bio varchar(255),
                profile_pic varchar(255),
                PRIMARY KEY (user_id)
);

CREATE TABLE calculations (
                calculation_id int,
                user_id int,
                payment_frequency varchar(10),
                salary float,
                tax float,
                "timestamp" timestamp,
                PRIMARY KEY (calculation_id),
                FOREIGN KEY (user_id) REFERENCES users (user_id)
)