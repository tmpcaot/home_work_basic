-- Индекс для ускорения поиска по пользователям
CREATE INDEX idx_users_email ON users(email);

-- Индекс для ускорения поиска по заказам пользователя
CREATE INDEX idx_orders_user_id ON orders(user_id);

-- Индекс для ускорения поиска по продуктам
CREATE INDEX idx_products_name ON products(name);
