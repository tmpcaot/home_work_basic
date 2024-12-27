-- Редактируем данные пользователя
UPDATE users
SET name = 'Иван Петров', email = 'ivan.petrov@example.com'
WHERE id = 1;

-- Редактируем данные товара
UPDATE products
SET name = 'Новый Продукт B', price = 75.00
WHERE id = 1;
