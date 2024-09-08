package repository

import (
    "context"
    "fmt"
    "bestlib-backend/models"
    "github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
    DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
    return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user *models.User) error {
    query := `
        INSERT INTO users (iin, full_name, password)
        VALUES ($1, $2, $3)
    `
    _, err := repo.DB.Exec(context.Background(), query, user.IIN, user.FullName, user.Password)
    if err != nil {
        return fmt.Errorf("error creating user: %w", err)
    }
    return nil
}

func (repo *UserRepository) FindUserByIIN(iin string) (*models.User, error) {
    var user models.User
    query := `
        SELECT id, iin, full_name, password
        FROM users
        WHERE iin = $1
    `
    row := repo.DB.QueryRow(context.Background(), query, iin)
    err := row.Scan(&user.ID, &user.IIN, &user.FullName, &user.Password)
    if err != nil {
        if err.Error() == "no rows in result set" {
            return nil, nil
        }
        return nil, fmt.Errorf("error finding user: %w", err)
    }
    return &user, nil
}

