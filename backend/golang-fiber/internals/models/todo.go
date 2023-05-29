package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
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
