package cmd

import (
	"context"
	"github.com/fatih/color"
	authapi "github.com/tylerjnettleton/go/pkg/auth/proto"
	"google.golang.org/grpc"
	"log"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login a user and get a JWT token back",
	Long:  `Usage example: login email@email.com password`,
	Run: func(cmd *cobra.Command, args []string) {
		callLoginFunction(args)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func callLoginFunction(args []string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := authapi.NewLoginClient(conn)
	response, err := c.LoginUser(context.Background(), &authapi.LoginMessage{Email: args[0], Password: args[1]})

	if err != nil {
		color.Red("Error when calling LoginUser: %s", err)
	}

	if response.Success {
		color.Green("Successfully logged in user: %s", args[0])
		color.Cyan("JWT Token: %s", response.Token)
	} else {
		color.Red("Invalid login for user %s", args[0])
	}
}
