CREATE TABLE IF NOT EXISTS customers(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    birth_date TIMESTAMP NOT NULL,
    gender VARCHAR(6) NOT NULL,
    email VARCHAR(55) NOT NULL,
    address VARCHAR(200) NULL
);

INSERT INTO customers(
    first_name,
    last_name,
    birth_date,
    gender,
    email,
    address) VALUES ('Liisa', 'Palusaar', '1999-06-22 19:10:33', 'Female', 'liisapalusaar@gmail.com', 'Kentmanni 3');

INSERT INTO customers(
    first_name,
    last_name,
    birth_date,
    gender,
    email,
    address) VALUES ('Roosa', 'Saare', '1996-01-22 14:20:33', 'Female', 'roosasaare@gmail.com', 'Estonia 4');

INSERT INTO customers(
    first_name,
    last_name,
    birth_date,
    gender,
    email,
    address) VALUES ('Ahmed', 'Abdullajev', '2000-01-04 14:20:55', 'Male', 'ahmed@gmail.com', 'Secret 00');