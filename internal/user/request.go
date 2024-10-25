package user

import (
	"time"
)

type RequestAddUsers struct {
	Name string `json:"name" binding:"required"`
}

type RequestGetUsers struct {
	Keyword string `json:"keyword"`
	Limit   *int   `json:"limit"`
	Offset  *int   `json:"offset"`
}

type Users struct {
	Id        string    `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Users) TableName() string {
	return "users"
}
