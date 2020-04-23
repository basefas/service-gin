package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"NOT NULL"`
	Password string `json:"password" gorm:"NOT NULL"`
	Email    string `json:"email"`
}

type CreateUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UpdateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type GetUserInfo struct {
	ID       uint   `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
