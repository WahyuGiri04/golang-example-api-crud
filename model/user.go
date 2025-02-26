package model

import "example-api/model/base"

type User struct {
	base.BaseModel
	Name     string `json:"name" binding:"required" validate:"required"`
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required"`
}

func (User) TableName() string {
	return "USER"
}
