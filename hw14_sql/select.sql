-- Выборка всех пользователей
SELECT * FROM users;

-- Выборка всех товаров
SELECT * FROM products;

-- Выборка заказов по пользователю
SELECT o.id, u.name, o.order_date, o.total_amount
FROM orders o
JOIN users u ON o.user_id = u.id
WHERE u.id = 1; 

-- Общая сумма заказов пользователя
SELECT SUM(o.total_amount) as total_sum
FROM orders o
WHERE o.user_id = 1;

-- Средняя цена товара у пользователя
SELECT AVG(p.price) as average_price
FROM order_products op
JOIN products p ON op.product_id = p.id
WHERE op.order_id IN (
    SELECT id
    FROM orders
    WHERE user_id = 1
); 
