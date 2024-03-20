-- Вставка пользователей
INSERT INTO Users (name, email, password) VALUES 
('Elena Belova', 'elena.belova@example.com', 'elenaSecure123'),
('Sergey Morozov', 'sergey.morozov@example.com', 'sergeyPass456'),
('Irina Zhukova', 'irina.zhukova@example.com', 'irina789Pass'),
('Nikolay Vasilyev', 'nikolay.vasilyev@example.com', 'nikolayPass234'),
('Olga Petrova', 'olga.petrova@example.com', 'olgaPetrova567'),
('Anna Kuznetsova', 'anna.kuznetsova@example.com', 'annaKuzPass890'),
('Vladimir Sokolov', 'vladimir.sokolov@example.com', 'vladimir321Pass');

-- Обновление информации о пользователях по id
UPDATE Users SET email = 'elena.newemail@example.com', password = 'newElenaPass123' WHERE id = 1;
UPDATE Users SET name = 'Sergey Novikov', email = 'sergey.novikov@example.com' WHERE id = 2;
UPDATE Users SET password = 'newIrinaPass789' WHERE id = 3;
UPDATE Users SET name = 'Nikolay Ivanov', password = 'nikolayNewPass234' WHERE id = 4;
UPDATE Users SET email = 'olga.newpetrova@example.com' WHERE id = 5;

-- Удаление пользователей по id
DELETE FROM Users WHERE id IN (6, 7);

-- Вставка товаров
INSERT INTO Products (name, price) VALUES 
('E-book', 15000.00),
('Coffee machine', 25000.00),
('Robot vacuum cleaner', 19000.00),
('Smartphone XL', 55000.00),
('Fitness bracelet', 7000.00),
('Wireless charging', 3000.00),
('External hard drive', 8000.00);

-- Обновление информации о товарах по id
UPDATE Products SET price = 16000.00 WHERE id = 1;
UPDATE Products SET name = 'Premium coffee machine', price = 30000.00 WHERE id = 2;
UPDATE Products SET name = 'Robot Vacuum Cleaner Smart', price = 21000.00 WHERE id = 3;
UPDATE Products SET price = 57000.00 WHERE id = 4;
UPDATE Products SET name = 'Advanced fitness bracelet', price = 9000.00 WHERE id = 5;

-- Удаление товаров по id
DELETE FROM Products WHERE id IN (6, 7);