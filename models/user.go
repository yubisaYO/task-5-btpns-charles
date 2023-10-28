package models

import (
    "time"
)

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string
    Email     string    `gorm:"unique"`
    Password  string
    CreatedAt time.Time
    UpdatedAt time.Time
    Photos    []Photo // Relasi dengan foto-foto yang dimiliki oleh pengguna
}
