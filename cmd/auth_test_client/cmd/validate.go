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

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a JWT token",
	Long:  `Example usage: validate JWTTOKEN`,
	Run: func(cmd *cobra.Command, args []string) {
		callValidateFuntion(args)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func callValidateFuntion(args []string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := authapi.NewValidateJWTTokenClient(conn)
	response, err := c.ValidateJWTToken(context.Background(), &authapi.ValidateJWTokenMessage{Token: args[0]})

	if err != nil {
		color.Red("Error when calling Validate Function: %s", err)
	}

	if response.Valid {
		color.Green("Token is valid!")
		if response.NewToken != "" {
			color.Cyan("Refreshed token: %s", response.NewToken)
		} else {
			color.Cyan("Refreshed token: None")
		}
	} else {
		color.Red("Invalid token!")
	}
}
