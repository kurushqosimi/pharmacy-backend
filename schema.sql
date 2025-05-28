-- Удаляем существующие таблицы (осторожно при проде)
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;

-- Таблица пользователей
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL
);

-- Таблица товаров
CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          description TEXT,
                          price NUMERIC(10, 2) NOT NULL,
                          image_url TEXT,
                          category VARCHAR(50),
                          stock INT DEFAULT 0
);

-- Таблица заказов
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INT REFERENCES users(id) ON DELETE CASCADE,
                        address TEXT NOT NULL,
                        phone VARCHAR(20) NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица позиций в заказе
CREATE TABLE order_items (
                             id SERIAL PRIMARY KEY,
                             order_id INT REFERENCES orders(id) ON DELETE CASCADE,
                             product_id INT REFERENCES products(id),
                             quantity INT NOT NULL CHECK (quantity > 0)
);

-- Индексы
CREATE INDEX idx_user_email ON users(email);
CREATE INDEX idx_product_name ON products(name);
CREATE INDEX idx_order_user ON orders(user_id);
CREATE INDEX idx_orderitem_order ON order_items(order_id);
