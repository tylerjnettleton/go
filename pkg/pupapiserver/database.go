package pupapiserver

import (
	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	authapi "github.com/tylerjnettleton/go/pkg/auth/proto"
	"google.golang.org/grpc"
	"log"
)

var db *gorm.DB

func CloseDB() {
	db.Close()
}

type RegisterError struct {
	Success         bool
	InvalidEmail    bool
	InvalidPassword bool
}

func Connect() error {
	// TODO: get login information from secure service
	tdb, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=test dbname=api password=test sslmode=disable")

	if err != nil {
		log.Print("Database connection failed")
		return err
	}

	db = tdb

	// Migrate users table
	db.AutoMigrate(&Profile{})

	return nil
}

func RegisterUser(user *User) (rsp *RegisterError) {
	// First register with the profile client
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5555", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := authapi.NewRegisterClient(conn)

	response, err := c.RegisterUser(context.Background(), &authapi.RegisterMessage{Email: user.Email, Password: user.Password})
	if err != nil {
		log.Print(err)
		return &RegisterError{
			Success: false,
		}
	}

	return &RegisterError{
		Success:         response.Success,
		InvalidPassword: response.InvalidPassword,
		InvalidEmail:    response.TakenEmail,
	}
}
