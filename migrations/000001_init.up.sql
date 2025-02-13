CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    balance INT DEFAULT 1000 NOT NULL
);

CREATE TABLE merchandise (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    price INT NOT NULL
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    sender TEXT REFERENCES users(username) ON DELETE CASCADE,
    receiver TEXT REFERENCES users(username) ON DELETE CASCADE,
    amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE purchases (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    merchandise_id INT REFERENCES merchandise(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT now()
);

INSERT INTO merchandise (name, price) VALUES 
    ('t-shirt', 80), ('cup', 20), ('book', 50), ('pen', 10), ('powerbank', 200),
    ('hoody', 300), ('umbrella', 200), ('socks', 10), ('wallet', 50), ('pink-hoody', 500);
