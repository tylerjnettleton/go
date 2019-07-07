package pupapiserver

import (
	"context"
	pupapi "github.com/tylerjnettleton/go/pkg/pupapiserver/proto"
)

type Server struct {
}

func (Server) RegisterUser(ctx context.Context, msg *pupapi.RegisterMessage) (*pupapi.RegisterResponse, error) {
	u := &User{
		FirstName: msg.FirstName,
		LastName:  msg.LastName,
		Email:     msg.Email,
		Password:  msg.Password,
		Zipcode:   msg.Zipcode,
	}

	err := RegisterUser(u)
	return &pupapi.RegisterResponse{
		Success:         err.Success,
		InvalidPassword: err.InvalidPassword,
		TakenEmail:      err.InvalidEmail,
	}, nil
}
