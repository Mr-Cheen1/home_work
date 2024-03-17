BEGIN;

-- Выборка статистики по пользователю: общая сумма заказов и средняя цена товара
SELECT 
    Users.id AS user_id,
    Users.name AS user_name,
    SUM(Orders.total_amount) AS total_order_amount, -- Общая сумма заказов пользователя
    ROUND(AVG(Products.price), 2) AS average_product_price -- Средняя цена товара в заказах пользователя с округлением до 2 знаков после запятой
FROM 
    Users
JOIN Orders ON Users.id = Orders.user_id
JOIN OrderProducts ON Orders.id = OrderProducts.order_id
JOIN Products ON OrderProducts.product_id = Products.id
GROUP BY Users.id
ORDER BY Users.id;

COMMIT;