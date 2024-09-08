package services

import (
    "bestlib-backend/repository"
    "bestlib-backend/utils"
)

type BookService struct {
    BookRepo *repository.BookRepository
}

func NewBookService(bookRepo *repository.BookRepository) *BookService {
    return &BookService{BookRepo: bookRepo}
}

func (svc *BookService) SearchBooks(searchTerm string) ([]map[string]interface{}, error) {
    bookInfo := utils.ExtractBookInfo(searchTerm)
    books, err := svc.BookRepo.SearchBooks(bookInfo.Title, bookInfo.Author, bookInfo.Year, bookInfo.Status)
    if err != nil {
        return nil, err
    }

    var bookResults []map[string]interface{}
    for _, book := range books {
        bookResults = append(bookResults, map[string]interface{}{
            "id":       book.ID,
            "author":   book.Author,
            "title":    book.Title,
            "year":     book.Year,
            "status":   book.Status,
        })
    }
    return bookResults, nil
}

