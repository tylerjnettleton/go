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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
		log.Printf("Successfully logged in user! %s", response.Token)
		color.Green("Successfully logged in user: %s", args[0])
		color.Cyan("JWT Token: %s", response.Token)
	} else {
		color.Red("Invalid login for user %s", args[0])
	}
}
