package main

import (
	"context"
	authapi "github.com/tylerjnettleton/pkg/auth/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	//c := authapi.NewRegisterClient(conn)
	//
	//response, err := c.RegisterUser(context.Background(), &authapi.RegisterMessage{Email: "tyler@tylernettleton.com", Password: "asdf"})
	//
	//if err != nil {
	//	log.Fatalf("Error when calling RegisterUser: %s", err)
	//}
	//if response.Success {
	//	log.Print("Successfully created user!")
	//} else {
	//	log.Printf("Failed to create user because - Email taken %t - Invalid Password %t", response.TakenEmail, response.InvalidPassword)
	//}

	c := authapi.NewLoginClient(conn)
	response, err := c.LoginUser(context.Background(), &authapi.LoginMessage{Email: "tyler@tylernettleton.com", Password: "asdf"})

	if err != nil {
		log.Fatalf("Error when calling LoginUser: %s", err)
	}
	if response.Success {
		log.Printf("Successfully logged in user! %s", response.Token)
	} else {
		log.Printf("Failed to login user")
	}
}
