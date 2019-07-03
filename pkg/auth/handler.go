package auth

import (
	"context"
	"github.com/tylerjnettleton/go/pkg/auth/proto"
	"log"
)

type Server struct {
}

func (s *Server) RegisterUser(ctx context.Context, req *authapi.RegisterMessage) (*authapi.RegisterResponse, error) {
	log.Printf("Receive message %s", req.Email)
	success, takenEmail, invalidPassword := RegisterUser(req.Email, req.Password)
	return &authapi.RegisterResponse{Success: success, TakenEmail: takenEmail, InvalidPassword: invalidPassword}, nil
}

func (s *Server) LoginUser(ctx context.Context, req *authapi.LoginMessage) (*authapi.LoginResponse, error) {
	err, token := AuthenticateUser(req.Email, req.Password)
	if err != nil {
		return &authapi.LoginResponse{Success: false, Token: ""}, nil
	}
	return &authapi.LoginResponse{Success: true, Token: token}, nil
}

func (s *Server) ValidateJWTToken(ctx context.Context, req *authapi.ValidateJWTokenMessage) (*authapi.ValidateJWTokenResponse, error) {
	valid, token := ValidateToken(req.Token)
	return &authapi.ValidateJWTokenResponse{Valid: valid, NewToken: token}, nil
}
