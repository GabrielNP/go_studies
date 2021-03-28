-- Copy, paste and run this scripts into a Postgres connection!

CREATE DATABASE web_store;

CREATE TABLE products (
    id serial primary key,
    name varchar,
    description varchar, 
    price decimal
    amount integer
);

INSERT INTO products (name, description, price, amount) VALUES
    ('T-shirt', 'Blue, very beaultiful', 39.9, 1),
    ('Sneakers', 'Comfortable', 60, 2),
    ('Headset', 'high-end', 200.54, 1);
