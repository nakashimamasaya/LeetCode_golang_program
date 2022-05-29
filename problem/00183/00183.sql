# Write your MySQL query statement below
SELECT c.name AS 'Customers'
FROM customers AS c  
where 1 > (
	SELECT COUNT(*)
	FROM orders AS o
	WHERE o.customerId = c.id
);
