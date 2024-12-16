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
