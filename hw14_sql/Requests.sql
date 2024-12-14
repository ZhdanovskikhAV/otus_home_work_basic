-- Вставка 4 пользователей
INSERT INTO AppAdmin.Users (name, email, password) VALUES 
('Иван Иванов', 'ivan@example.com', 'secure_password1'),
('Анна Смирнова', 'anna@example.com', 'secure_password2'),
('Петр Петров', 'petr@example.com', 'secure_password3'),
('Мария Кузнецова', 'maria@example.com', 'secure_password4');

-- Обновление 1 пользователя (например, с ID 1)
UPDATE AppAdmin.Users 
SET name = 'Иван Петров', email = 'ivan.petrov@example.com' 
WHERE id = 1;

-- Удаление 1 пользователя (например, с ID 2)
DELETE FROM AppAdmin.Users 
WHERE id = 2;

-- Выборка всех пользователей
SELECT * FROM AppAdmin.Users;

-- Вставка 4 продукта
INSERT INTO AppAdmin.Products (name, price) VALUES 
('Товар 1', 100.00),
('Товар 2', 150.00),
('Товар 3', 200.00),
('Товар 4', 250.00);

-- Обновление 1 продукта (например, с ID 1)
UPDATE AppAdmin.Products 
SET name = 'Обновленный Товар 1', price = 120.00 
WHERE id = 1;

-- Удаление 1 продукта (например, с ID 3)
DELETE FROM AppAdmin.Products 
WHERE id = 3;

-- Выборка всех продуктов
SELECT * FROM AppAdmin.Products;

-- Вставка данных в таблицу Orders
INSERT INTO AppAdmin.Orders (user_id, order_date, total_amount) VALUES 
(1, '2023-10-01 10:00:00', 250.00),  -- Заказ от пользователя с ID 1
(3, '2023-10-02 11:30:00', 450.00);  -- Заказ от пользователя с ID 3

-- Вставка данных в таблицу OrderProducts
INSERT INTO AppAdmin.OrderProducts (order_id, product_id, quantity) VALUES 
(5, 1, 2),  -- 2 штуки товара с ID 1 в заказе с ID 1
(5, 2, 1),  -- 1 штука товара с ID 2 в заказе с ID 1
(6, 2, 3),  -- 3 штуки товара с ID 3 в заказе с ID 2
(6, 4, 1);  -- 1 штука товара с ID 4 в заказе с ID 2

-- Удаление заказа с ID 1 и всех связанных с ним продуктов
--DELETE FROM AppAdmin.OrderProducts 
--WHERE order_id = 1;  -- Удаляем все продукты из заказа с ID 1

--DELETE FROM AppAdmin.Orders 
--WHERE id = 1;  -- Удаляем сам заказ с ID 1

-- Выборка всех оставшихся заказов
SELECT * FROM AppAdmin.Orders;

-- Выборка всех оставшихся продуктов по заказам
SELECT * FROM AppAdmin.OrderProducts;

-- Выборка заказов по пользователю с user_id = 1
SELECT * 
FROM AppAdmin.Orders 
WHERE user_id = 1; 

-- Выборка общая сумма заказов/средняя цена товара по пользователю 
SELECT 
    o.user_id,
    SUM(o.total_amount) AS total_order_amount,  
    AVG(p.price) AS average_product_price        
FROM 
    AppAdmin.Orders o
INNER JOIN 
    AppAdmin.OrderProducts op ON o.id = op.order_id
INNER JOIN 
    AppAdmin.Products p ON op.product_id = p.id
WHERE 
    o.user_id = 1  
GROUP BY 
    o.user_id;

