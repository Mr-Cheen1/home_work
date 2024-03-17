BEGIN;

-- Выборка заказов для всех пользователей с именами
SELECT 
    Orders.user_id,
    Users.name AS user_name,
    Orders.id AS order_id,
    Orders.order_date,
    Orders.total_amount,
    COUNT(OrderProducts.product_id) AS total_products, -- Общее количество различных товаров в заказе
    SUM(OrderProducts.quantity) AS total_quantity -- Общее количество товаров в заказе
FROM 
    Orders
LEFT JOIN OrderProducts ON Orders.id = OrderProducts.order_id
LEFT JOIN Users ON Orders.user_id = Users.id 
GROUP BY Orders.id, Users.name, Orders.user_id 
ORDER BY Orders.user_id, Orders.order_date DESC;

COMMIT;