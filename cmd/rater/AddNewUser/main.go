package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
	"os"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

type book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func hello(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	bk := &book{
		ISBN:   "978-1420931693",
		Title:  "The Republic",
		Author: "Plato",
	}

	js, err := json.Marshal(bk)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}


// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func main() {
	lambda.Start(hello)
}