-- Вставка данных в таблицу Orders
INSERT INTO AppAdmin.Orders (user_id, order_date, total_amount) VALUES 
(1, '2023-10-01 10:00:00', 250.00),  -- Заказ от пользователя с ID 1
(2, '2023-10-02 11:30:00', 450.00);  -- Заказ от пользователя с ID 2

-- Вставка данных в таблицу OrderProducts
INSERT INTO AppAdmin.OrderProducts (order_id, product_id, quantity) VALUES 
(1, 1, 2),  -- 2 штуки товара с ID 1 в заказе с ID 1
(1, 2, 1),  -- 1 штука товара с ID 2 в заказе с ID 1
(2, 3, 3),  -- 3 штуки товара с ID 3 в заказе с ID 2
(2, 4, 1);  -- 1 штука товара с ID 4 в заказе с ID 2

-- Удаление заказа с ID 1 и всех связанных с ним продуктов
DELETE FROM AppAdmin.OrderProducts 
WHERE order_id = 1;  -- Удаляем все продукты из заказа с ID 1

DELETE FROM AppAdmin.Orders 
WHERE id = 1;  -- Удаляем сам заказ с ID 1

-- Выборка всех оставшихся заказов
SELECT * FROM AppAdmin.Orders;

-- Выборка всех оставшихся продуктов по заказам
SELECT * FROM AppAdmin.OrderProducts;

-- Выборка заказов по пользователю с user_id = 1
SELECT * 
FROM AppAdmin.Orders 
WHERE user_id = 1; 
