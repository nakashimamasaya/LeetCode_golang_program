
-- +migrate Up
CREATE TABLE customers(
	id int PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(255)
);

CREATE TABLE orders(
	id INT PRIMARY KEY AUTO_INCREMENT,
	customerId INT,
	INDEX customer_index(customerId),
	FOREIGN KEY fk_customer(customerId) REFERENCES customers(id)
);

-- +migrate Down
DROP TABLE orders;
DROP TABLE customers;
