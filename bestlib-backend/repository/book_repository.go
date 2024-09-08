package repository

import (
    "context"
    "bestlib-backend/models"
    "github.com/jackc/pgx/v5/pgxpool"
)

type BookRepository struct {
    DB *pgxpool.Pool
}

func NewBookRepository(db *pgxpool.Pool) *BookRepository {
    return &BookRepository{DB: db}
}

func (repo *BookRepository) SearchBooks(title, author, year, status string) ([]models.Book, error) {
    var books []models.Book

    query := `
        SELECT id, title, author, year, status
        FROM books
        WHERE title ILIKE $1 AND author ILIKE $2 AND year ILIKE $3 AND status ILIKE $4
    `
    rows, err := repo.DB.Query(context.Background(), query, "%"+title+"%", "%"+author+"%", "%"+year+"%", "%"+status+"%")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var book models.Book
        if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.Status); err != nil {
            return nil, err
        }
        books = append(books, book)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return books, nil
}
