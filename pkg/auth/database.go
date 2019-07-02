package auth

import (
	"encoding/base64"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func CloseDB() {
	db.Close()
}

func Connect() error {
	// TODO: get login information from secure service
	db, _ = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=test dbname=user_auth password=test sslmode=disable")

	//if err != nil {
	//	return err
	//}

	// Migrate users table
	db.AutoMigrate(&User{})

	// TODO: Maybe put this into a context to close the database on application termination?
	return nil
}

func RegisterUser(email string, password string) (success bool, takenEmail bool, invalidPassword bool) {
	// Check if user with the same email exist
	user := &User{}
	db.Where(&User{Email: email}).First(user)

	if user.Email != "" {
		return false, true, false
	}

	// Email does not exist, create a new user after hashing password
	salt, _ := randbytes()
	hashedPassword, _ := HashPassword([]byte(password), salt)

	newUser := &User{
		Email:          email,
		HashedPassword: base64.StdEncoding.EncodeToString(hashedPassword),
		Salt:           base64.StdEncoding.EncodeToString(salt),
	}

	db.Create(&newUser)

	//TODO: check the password constrains
	return true, false, false
}

func AuthenticateUser(email string, password string) (err error, token string) {

	var user User
	res := db.Find(&user, &User{Email: email})

	if res.Error != nil {
		// Todo: return an actual, useful error
		return res.Error, ""
	}

	// Attempt to authenticate the user
	decodedHashedPassword, _ := base64.StdEncoding.DecodeString(user.HashedPassword)
	decodedSalt, _ := base64.StdEncoding.DecodeString(user.Salt)

	authenticated, err := Authenticate([]byte(password), decodedSalt, decodedHashedPassword)

	if authenticated {
		return GenerateTokenForUser(user)
	} else {
		return errors.New("invalid login provided"), ""
	}
}
