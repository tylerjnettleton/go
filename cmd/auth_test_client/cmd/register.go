/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	authapi "github.com/tylerjnettleton/go/pkg/auth/proto"
	"google.golang.org/grpc"
	"log"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a user [email, password)",
	Long:  `Usage example: login email@email.com password`,
	Run: func(cmd *cobra.Command, args []string) {
		callRegisterFunction(args)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}

func callRegisterFunction(args []string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := authapi.NewRegisterClient(conn)

	response, err := c.RegisterUser(context.Background(), &authapi.RegisterMessage{Email: args[0], Password: args[1]})

	if err != nil {
		log.Fatalf("Error when calling RegisterUser: %s", err)
	}
	if response.Success {
		color.Green("Success! Registered a new user (%s)", args[0])
	} else {

		color.Red("Failed! Could not register a new user (%s)", args[0])
		if response.TakenEmail {
			color.Blue("*** Reason ***")
			color.Cyan("1.) Email was already taken")
		}

		if response.InvalidPassword {
			color.Blue("*** Reason ***")
			color.Cyan("1.) Invalid password")
		}
	}
}
