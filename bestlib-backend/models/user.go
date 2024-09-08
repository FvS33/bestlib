package models

type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    IIN      string `json:"iin" gorm:"unique;not null"`
    FullName string `json:"full_name" gorm:"not null"`
    Password string `json:"password" gorm:"not null"`
}
