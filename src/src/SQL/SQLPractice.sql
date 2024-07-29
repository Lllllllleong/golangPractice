--
-- SQL Practice
--

select name, population, area
from World
where area >= 3000000 or population >= 25000000


select product_id from Products
where low_fats = 'Y' && recyclable = 'Y';


select name from customer
where referee_id is null or referee_id != 2;


select
    employee_id,
    IF (employee_id % 2 == 1 && name not regexp '^M', salary, 0) as bonus
from
    employees
order by
    employee_id;



