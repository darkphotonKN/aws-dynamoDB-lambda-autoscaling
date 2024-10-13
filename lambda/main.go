package main

import (
	"fmt"
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Username string `json:"username"`
}

func HandleEvent(event Event) (string, error) {

	if event.Username == "" {

		return "", fmt.Errorf("username in event %v cannot be empty", event)
	}

	return fmt.Sprintf("Sucessfuly called HandleEvent -%s", event.Username), nil
}

func main() {
	myApp := app.NewApp()
	lambda.Start(myApp.ApiHandler.RegisterUserHandler)
}
