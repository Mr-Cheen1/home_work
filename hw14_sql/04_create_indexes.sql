BEGIN;

-- Индексы для таблицы Users
CREATE INDEX idx_users_name ON Users(name);
CREATE INDEX idx_users_email ON Users(email);

-- Индексы для таблицы Orders
CREATE INDEX idx_orders_user_id ON Orders(user_id);
CREATE INDEX idx_orders_order_date ON Orders(order_date);
CREATE INDEX idx_orders_total_amount ON Orders(total_amount);

-- Индексы для таблицы Products
CREATE INDEX idx_products_name ON Products(name);
CREATE INDEX idx_products_price ON Products(price);

-- Индексы для таблицы OrderProducts
CREATE INDEX idx_orderproducts_order_id ON OrderProducts(order_id);
CREATE INDEX idx_orderproducts_product_id ON OrderProducts(product_id);

COMMIT;