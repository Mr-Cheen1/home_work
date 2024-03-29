BEGIN;

-- Таблица "Пользователи" (Users) - содержит информацию о пользователях
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Таблица "Заказы" (Orders) - отображает информацию о заказах
CREATE TABLE Orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    order_date DATE NOT NULL,
    total_amount DECIMAL NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

-- Таблица "Товары" (Products) - содержит информацию о товарах
CREATE TABLE Products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL NOT NULL
);

-- Таблица "Заказы-Товары" (OrderProducts) - содержит информацию о отношении заказов к товарам (многие ко многим)
CREATE TABLE OrderProducts (
    order_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    FOREIGN KEY (order_id) REFERENCES Orders(id),
    FOREIGN KEY (product_id) REFERENCES Products(id),
    PRIMARY KEY (order_id, product_id)
);

COMMIT;