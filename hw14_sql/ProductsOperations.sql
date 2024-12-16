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