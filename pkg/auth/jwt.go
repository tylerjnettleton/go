package auth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// TODO: Replace this with a secret key
var jwtKey = []byte("my_secret_key")

type jwtToken struct {
	Email  string `json:"username"`
	UserId uint
	jwt.StandardClaims
}

func GenerateTokenForUser(user User) (err error, token string) {
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	jt := &jwtToken{
		Email:  user.Email,
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jt)
	// Create the JWT string
	tString, err := t.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return err, ""
	}

	return nil, tString
}

func ValidateToken(tknStr string) (valid bool, tokenString string) {
	jt := &jwtToken{}

	// Parse the JWT string and store the result
	tkn, err := jwt.ParseWithClaims(tknStr, jt, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		log.Printf("Invalid token for user: %s", jt.Email)
		return false, ""
	}

	if err != nil {
		log.Printf("Invalid token for user: %s", jt.Email)
		return false, ""
	}

	log.Printf("Valid token for user: %s", jt.Email)
	return true, RefreshToken(tknStr)
}

func RefreshToken(tknStr string) (tString string) {
	jt := &jwtToken{}

	// Parse the JWT string and store the result
	tkn, err := jwt.ParseWithClaims(tknStr, jt, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return ""
	}

	if !tkn.Valid {
		return ""
	}

	// Extend the time on the token

	if time.Unix(jt.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		// No need to extend the excoriation time
		// return the old one
		// (This should not really happen)
		return ""
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	jt.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jt)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return tknStr
	}

	return tokenString
}
