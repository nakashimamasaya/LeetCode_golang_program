SELECT e_one.`name` as "Employee"
FROM employee as e_one
WHERE e_one.managerId IS NOT NULL
AND e_one.salary > (
	SELECT e_two.salary
	FROM employee as e_two
	WHERE e_two.id = e_one.managerId  
);
