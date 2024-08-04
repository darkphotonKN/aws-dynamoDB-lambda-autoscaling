package main

import (
	"fmt"
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
	lambda.Start(HandleEvent)
}
