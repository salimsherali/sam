USE sneakfit_sam;

--password 123456
INSERT INTO users (name, email, password, role)
VALUES (
    'Admin User',
    'admin@admin.com',
    '$2a$10$7QJ8z9YlXQvYl8X7wFJcUuQy6k9rV5j3z1vZzX9YlXQvYl8X7wFJcU', 
    'admin'
);
