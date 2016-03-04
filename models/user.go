package models

import (
    "time"
    "github.com/satori/go.uuid"
)

type User struct {
    Id          string      `gorm:"primary_key:true"`
    Name        string                                  `sql:"not null"`
    Email       string                                  `sql:"unique_index;not null"`
    Password    string                                  `sql:"not null"`
    CreatedAt   time.Time   `gorm:"column:Created"`
    UpdatedAt   time.Time   `gorm:"column:Updated"`
}

func (this *User) BeforeCreate() {
    if this.Id == "" { this.Id = uuid.NewV4().String() }
}

// usuario default da nossa app
var UserDefault = &User {
    Id:"75f0a8d6-736f-4d8a-8ff4-e37573a0a516",
    Name:"John Doe",
    Email:"my@email.com",
    Password:"workshop",
}