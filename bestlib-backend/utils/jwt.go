package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
    "errors"
)

var secretKey = []byte("your_secret_key")

func GenerateJWT(iin string) (string, error) {
    claims := &jwt.StandardClaims{
        Subject: iin,
        ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

func ParseJWT(tokenString string) (*jwt.StandardClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, errors.New("invalid token")
}

