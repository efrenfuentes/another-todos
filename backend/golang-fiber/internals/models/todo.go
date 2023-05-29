package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	Url       string `json:"url"`
}

type UpdateTodo struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
	Order     *int    `json:"order"`
	Url       *string `json:"url"`
}
