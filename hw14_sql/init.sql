-- 1. Создание пользователя AppAdmin с полными правами
CREATE ROLE AppAdmin WITH LOGIN PASSWORD '123456';
ALTER ROLE AppAdmin CREATEDB;  -- Позволяет пользователю создавать базы данных

-- 2. Создание базы данных STORE
CREATE DATABASE STORE OWNER AppAdmin;

-- 3. Назначение прав на базу данных STORE для AppAdmin
GRANT ALL PRIVILEGES ON DATABASE STORE TO AppAdmin;

-- 4. Подключение к базе данных STORE
--\c STORE;

-- 5. Создание схемы AppAdmin в базе данных STORE
CREATE SCHEMA IF NOT EXISTS AppAdmin;

-- 6. Создание таблицы Пользователи (Users) в схеме AppAdmin
CREATE TABLE IF NOT EXISTS AppAdmin.Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Создание индекса для ускорения выборки по email
CREATE INDEX idx_users_email ON AppAdmin.Users(email);

-- 7. Создание таблицы Заказы (Orders) в схеме AppAdmin
CREATE TABLE IF NOT EXISTS AppAdmin.Orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES AppAdmin.Users(id) ON DELETE CASCADE,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_amount DECIMAL(10, 2) NOT NULL
);

-- Создание индекса для ускорения выборки по user_id
CREATE INDEX idx_orders_user_id ON AppAdmin.Orders(user_id);

-- 8. Создание таблицы Товары (Products) в схеме AppAdmin
CREATE TABLE IF NOT EXISTS AppAdmin.Products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

-- Создание индекса для ускорения выборки по name
CREATE INDEX idx_products_name ON AppAdmin.Products(name);

-- 9. Создание таблицы Заказы-Товары (OrderProducts) в схеме AppAdmin
CREATE TABLE IF NOT EXISTS AppAdmin.OrderProducts (
    order_id INT REFERENCES AppAdmin.Orders(id) ON DELETE CASCADE,
    product_id INT REFERENCES AppAdmin.Products(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    PRIMARY KEY (order_id, product_id)
);

-- Создание индексов для ускорения выборки
CREATE INDEX idx_orderproducts_order_id ON AppAdmin.OrderProducts(order_id);
CREATE INDEX idx_orderproducts_product_id ON AppAdmin.OrderProducts(product_id);
