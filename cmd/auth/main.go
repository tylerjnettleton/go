package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/tylerjnettleton/go/pkg/auth"
	authapi "github.com/tylerjnettleton/go/pkg/auth/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Start monitoring for OS termination signals
	go monitorOSSignals()

	// Connect to database
	err := auth.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database %s", err)
	}

	// TODO: use TLS...
	// Create grpc listen handle
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create the GRPC auth.Server
	s := auth.Server{}
	grpcServer := grpc.NewServer()

	// Register GRPC functions
	authapi.RegisterRegisterServer(grpcServer, &s)
	authapi.RegisterLoginServer(grpcServer, &s)
	authapi.RegisterValidateJWTTokenServer(grpcServer, &s)

	color.Green("Starting Authentication Server...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server %s", err)
	}
}

func monitorOSSignals() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	sig := <-gracefulStop
	auth.CloseDB()
	log.Printf("Terminating auth server because of sig: %+v", sig)
	log.Print("Wait for 2 second to finish processing")
	time.Sleep(2 * time.Second)
	os.Exit(0)
}
