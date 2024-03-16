BEGIN;

-- Выборка информации о пользователях
-- Выбираем пользователей, добавляем информацию о количестве их заказов и общей сумме заказов, сортируем по количеству заказов
SELECT 
    Users.id, 
    Users.name, 
    Users.email, 
    COUNT(Orders.id) AS order_count, -- Количество заказов каждого пользователя
    SUM(Orders.total_amount) AS total_spent -- Общая сумма заказов каждого пользователя
FROM 
    Users
LEFT JOIN Orders ON Users.id = Orders.user_id -- LEFT JOIN, чтобы включить пользователей без заказов
GROUP BY Users.id
ORDER BY order_count DESC, total_spent DESC; -- Сортировка по количеству заказов и общей сумме

-- Выборка информации о товарах
-- Выбираем товары, добавляем информацию о количестве раз, когда товар был заказан, и общем количестве заказанных единиц, сортируем по популярности
SELECT 
    Products.id, 
    Products.name, 
    Products.price, 
    COUNT(OrderProducts.product_id) AS times_ordered, -- Количество раз, когда товар был заказан
    SUM(OrderProducts.quantity) AS total_quantity_ordered -- Общее количество заказанных единиц
FROM 
    Products
LEFT JOIN OrderProducts ON Products.id = OrderProducts.product_id -- LEFT JOIN, чтобы включить товары, которые не были заказаны
GROUP BY Products.id
ORDER BY times_ordered DESC, total_quantity_ordered DESC; -- Сортировка по популярности товара

COMMIT;