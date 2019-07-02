package auth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// TODO: Replace this with a secret key
var jwtKey = []byte("my_secret_key")

type jwtToken struct {
	Email string `json:"username"`
	jwt.StandardClaims
}

func GenerateTokenForUser(user User) (err error, token string) {
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	jt := &jwtToken{
		Email: user.Email,
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

func ValidateToken(token string) (valid bool) {
	jt := &jwtToken{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, jt, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		log.Fatalf("Invalid token for user: %s", jt.Email)
		return false
	}

	if err != nil {
		log.Fatalf("Invalid token for user: %s", jt.Email)
		return false
	}

	log.Printf("Valid token for user: %s", jt.Email)
	return true
}
