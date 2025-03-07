\c bestlib

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    iin VARCHAR(20) UNIQUE NOT NULL,
    fullname VARCHAR(100) NOT NULL,
    password TEXT NOT NULL,
);

CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    year INT,
    status VARCHAR(20) DEFAULT 'available'
);

INSERT INTO books (title, author, year, status) VALUES
('To Kill a Mockingbird', 'Harper Lee', 1960, 'available'),
('1984', 'George Orwell', 1949, 'available'),
('The Great Gatsby', 'F. Scott Fitzgerald', 1925, 'available'),
('Pride and Prejudice', 'Jane Austen', 1813, 'available'),
('The Catcher in the Rye', 'J.D. Salinger', 1951, 'available'),
('The Hobbit', 'J.R.R. Tolkien', 1937, 'available'),
('Fahrenheit 451', 'Ray Bradbury', 1953, 'available'),
('Brave New World', 'Aldous Huxley', 1932, 'available'),
('Moby-Dick', 'Herman Melville', 1851, 'available'),
('War and Peace', 'Leo Tolstoy', 1869, 'available');
