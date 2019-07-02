package auth

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email          string `json:”email”`
	HashedPassword string `json:”hashed_password”`
	Salt           string `json:”salt”`
}
