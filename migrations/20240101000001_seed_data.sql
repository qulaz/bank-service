-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- Тестовый пользователь (пароль: password)
INSERT INTO users (id, username, password, email, full_name, created_at, updated_at)
VALUES (1, 'testuser', '$2a$10$1XOzMzVkVDZVA9Vk6UbgL.6AEGDn7CUC9fVpUz9RqMGtWx9tTdkNi', 'test@example.com', 'Test User', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO users (id, username, password, email, full_name, created_at, updated_at)
VALUES (2, 'testuser2', '$2a$10$1XOzMzVkVDZVA9Vk6UbgL.6AEGDn7CUC9fVpUz9RqMGtWx9tTdkNi', 'test2@example.com', 'Test User 2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Тестовые счета
INSERT INTO accounts (id, user_id, number, balance, currency, created_at, updated_at)
VALUES (1, 1, '1234567890123456', 100000, 'RUB', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO accounts (id, user_id, number, balance, currency, created_at, updated_at)
VALUES (2, 1, '1234867549086794', 50433, 'RUB', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO accounts (id, user_id, number, balance, currency, created_at, updated_at)
VALUES (3, 1, '6543210987654321', 50000, 'USD', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO accounts (id, user_id, number, balance, currency, created_at, updated_at)
VALUES (4, 2, '1293829428954734', 10344200, 'RUB', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO accounts (id, user_id, number, balance, currency, created_at, updated_at)
VALUES (5, 2, '8946358459348594', 2374, 'USD', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO otps (user_id, code, expires_at, used, created_at)
VALUES 
(1, 530439, datetime('now', '-59 minutes'), FALSE, datetime('now')),
(1, 143956, datetime('now', '-58 minutes'), FALSE, datetime('now')),
(1, 985424, datetime('now', '-57 minutes'), FALSE, datetime('now')),
(1, 865923, datetime('now', '-56 minutes'), FALSE, datetime('now')),
(1, 875459, datetime('now', '-55 minutes'), FALSE, datetime('now')),
(1, 659356, datetime('now', '-54 minutes'), FALSE, datetime('now')),
(1, 986494, datetime('now', '-53 minutes'), FALSE, datetime('now')),
(1, 567542, datetime('now', '-52 minutes'), FALSE, datetime('now')),
(1, 123456, datetime('now', '-51 minutes'), FALSE, datetime('now')),
(1, 234567, datetime('now', '-50 minutes'), FALSE, datetime('now')),
(1, 345678, datetime('now', '-45 minutes'), FALSE, datetime('now')),
(1, 456789, datetime('now', '-40 minutes'), FALSE, datetime('now')),
(1, 567890, datetime('now', '-35 minutes'), FALSE, datetime('now')),
(1, 678901, datetime('now', '-30 minutes'), FALSE, datetime('now')),
(1, 789012, datetime('now', '-25 minutes'), FALSE, datetime('now')),
(1, 890123, datetime('now', '-20 minutes'), FALSE, datetime('now')),
(1, 901234, datetime('now', '-15 minutes'), FALSE, datetime('now')),
(1, 112233, datetime('now', '-10 minutes'), FALSE, datetime('now')),
(1, 223344, datetime('now', '-9 minutes'), FALSE, datetime('now')),
(1, 334455, datetime('now', '-8 minutes'), FALSE, datetime('now')),
(1, 445566, datetime('now', '-7 minutes'), FALSE, datetime('now')),
(1, 556677, datetime('now', '-6 minutes'), FALSE, datetime('now')),
(1, 667788, datetime('now', '-5 minutes'), FALSE, datetime('now')),
(1, 778899, datetime('now', '-4 minutes'), FALSE, datetime('now')),
(1, 889900, datetime('now', '-3 minutes'), FALSE, datetime('now')),
(1, 990011, datetime('now', '-2 minutes'), FALSE, datetime('now')),
(1, 100111, datetime('now', '-1 minute'), FALSE, datetime('now')),
(1, 110022, datetime('now'), FALSE, datetime('now')),
(1, 22033, datetime('now', '+1 minute'), FALSE, datetime('now')),
(1, 330044, datetime('now', '-45 minutes'), FALSE, datetime('now')),
(1, 440055, datetime('now', '-40 minutes'), FALSE, datetime('now')),
(1, 550066, datetime('now', '-80 minutes'), FALSE, datetime('now')),
(1, 660077, datetime('now', '-75 minutes'), FALSE, datetime('now')),
(1, 770088, datetime('now', '-70 minutes'), FALSE, datetime('now')),
(1, 880099, datetime('now', '-65 minutes'), FALSE, datetime('now')),
(1, 990000, datetime('now', '-60 minutes'), FALSE, datetime('now')),
(1, 111222, datetime('now', '-55 minutes'), FALSE, datetime('now')),
(1, 222333, datetime('now', '-50 minutes'), FALSE, datetime('now')),
(1, 333444, datetime('now', '-45 minutes'), FALSE, datetime('now')),
(1, 444555, datetime('now', '-40 minutes'), FALSE, datetime('now')),
(1, 555666, datetime('now', '-35 minutes'), FALSE, datetime('now')),
(1, 666777, datetime('now', '-30 minutes'), FALSE, datetime('now')),
(1, 777888, datetime('now', '-25 minutes'), FALSE, datetime('now')),
(1, 88999, datetime('now', '-20 minutes'), FALSE, datetime('now')),
(1, 999000, datetime('now', '-15 minutes'), FALSE, datetime('now')),
(1, 112233, datetime('now', '-10 minutes'), FALSE, datetime('now')),
(1, 223344, datetime('now', '-9 minutes'), FALSE, datetime('now')),
(1, 334455, datetime('now', '-8 minutes'), FALSE, datetime('now')),
(1, 445566, datetime('now', '-7 minutes'), FALSE, datetime('now')),
(1, 556677, datetime('now', '-6 minutes'), FALSE, datetime('now')),
(1, 667788, datetime('now', '-5 minutes'), FALSE, datetime('now')),
(1, 77899, datetime('now', '-4 minutes'), FALSE, datetime('now')),
(1, 889900, datetime('now', '-3 minutes'), FALSE, datetime('now')),
(1, 990011, datetime('now', '-2 minutes'), FALSE, datetime('now'));


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DELETE FROM accounts WHERE user_id = 1;
DELETE FROM users WHERE username = 'testuser'; 