-- Створення таблиці користувачів
CREATE TABLE users (
    id SERIAL PRIMARY KEY, -- Унікальний ідентифікатор
    username VARCHAR(50) UNIQUE NOT NULL, -- Ім'я користувача
    email VARCHAR(100) UNIQUE NOT NULL, -- Email
    password_hash VARCHAR(255) NOT NULL, -- Хеш пароля
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Час створення
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Час останнього оновлення
);

-- Створення таблиці публікацій
CREATE TABLE posts (
    id SERIAL PRIMARY KEY, -- Унікальний ідентифікатор публікації
    user_id INTEGER NOT NULL, -- Ідентифікатор користувача
    title VARCHAR(100) NOT NULL, -- Заголовок публікації
    content TEXT NOT NULL, -- Зміст публікації
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Час створення
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Час останнього оновлення
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE -- Зовнішній ключ
);

-- Індекси для покращення продуктивності
CREATE INDEX idx_users_username ON users(username); -- Індекс на username
CREATE INDEX idx_posts_user_id ON posts(user_id); -- Індекс на user_id
CREATE INDEX idx_posts_title ON posts(title); -- Індекс на заголовок
