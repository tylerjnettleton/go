package pupapiserver

import "github.com/jinzhu/gorm"

type User struct {
	Email     string `json:”email”`
	FirstName string `json:”first_name”`
	LastName  string `json:”last_name”`
	Zipcode   string `json:”zipcode”`
	Password  string `json:”password”`
}
type Profile struct {
	gorm.Model
	Email     string `json:”email”`
	FirstName string `json:”first_name”`
	LastName  string `json:”last_name”`
	Zipcode   string `json:”zipcode”`
}
