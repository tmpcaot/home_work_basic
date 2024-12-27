-- Вставляем пользователя
INSERT INTO users (name, email, password)
VALUES ('Иван Иванов', 'ivan@example.com', '12345');

-- Вставляем товар
INSERT INTO products (name, price)
VALUES ('Продукт A', 50.00);

-- Сохраняем заказ
WITH inserted_order AS (
    INSERT INTO orders (user_id, total_amount)
    VALUES (1, 150.00)
    RETURNING id
)
INSERT INTO order_products (order_id, product_id, quantity)
SELECT id, 1, 5
FROM inserted_order;
