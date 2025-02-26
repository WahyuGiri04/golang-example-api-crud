package model

import "example-api/model/base"

type Role struct {
	base.BaseModel
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (Role) TableName() string {
	return "ROLE"
}
