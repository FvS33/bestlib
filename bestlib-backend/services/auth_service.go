package services

import (
    "bestlib-backend/models"
    "bestlib-backend/repository"
    "bestlib-backend/utils"
    "golang.org/x/crypto/bcrypt"
    "errors"
)

type AuthService struct {
    UserRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
    return &AuthService{UserRepo: userRepo}
}

func (svc *AuthService) RegisterUser(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return svc.UserRepo.CreateUser(user)
}

func (svc *AuthService) LoginUser(iin, password string) (string, error) {
    user, err := svc.UserRepo.FindUserByIIN(iin)
    if err != nil {
        return "", err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    token, err := utils.GenerateJWT(iin)
    if err != nil {
        return "", err
    }
    return token, nil
}

