package main

import (
	"fmt"
	"lambda-func/app"
	"lambda-func/middleware"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Username string `json:"username"`
}

func HandleRequest(event Event) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by - %s", event.Username), nil
}

func ProtectedHandler(reqest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body: "This is a secret path",
		StatusCode: http.StatusOK,
	}, nil }

func main() {
	newApp := app.NewApp()
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.Path {
		case "/register":
			return newApp.ApiHandler.RegisterUserHandler(request)
		case "/login":
			return newApp.ApiHandler.LoginUser(request)
		case "/protected":
			return middleware.ValidateJwtMiddleware(ProtectedHandler)(request)
		default:
			return events.APIGatewayProxyResponse{
				Body: "Not found",
				StatusCode: http.StatusNotFound,
			}, nil
	}
})
}