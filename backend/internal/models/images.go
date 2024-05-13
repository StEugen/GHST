package models

import (
    "gorm.io/gorm"
)

type Image struct {
    gorm.Model
    Name      string
    Tag       string
    Digest    string
    Size      int64
    UserID    uint
    User      User
}

