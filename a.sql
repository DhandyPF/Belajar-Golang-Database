CREATE TABLE customer (
    id VARCHAR(10) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

SELECT * FROM customer;

DELETE FROM customer;

ALTER TABLE customer
ADD COLUMN email VARCHAR(100),
ADD COLUMN balance INT DEFAULT 0,
ADD COLUMN rating FLOAT DEFAULT 0.0, 
ADD COLUMN birth_date DATE, 
ADD COLUMN married BOOLEAN DEFAULT FALSE, 
ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

INSERT INTO customer(id, name, email, balance, rating, birth_date, married, created_at) VALUES('1', 'Dhandy', 'dhandy@gmail.com', 1000, 5.0, '2007-03-10', FALSE, CURRENT_TIMESTAMP);
INSERT INTO customer(id, name, email, balance, rating, birth_date, married, created_at) VALUES('2', 'Dimas', 'dimas@gmail.com', 1000, 5.0, '2007-03-15', TRUE, CURRENT_TIMESTAMP);
INSERT INTO customer(id, name, balance, rating, married, created_at) VALUES('3', 'Joko', 1000, 5.0, TRUE, CURRENT_TIMESTAMP);

UPDATE customer SET balance = 666666 WHERE id = '1';
UPDATE customer SET balance = 50000 WHERE id = '2';
UPDATE customer SET balance = 75000 WHERE id = '3';