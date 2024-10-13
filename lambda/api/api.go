package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

// request to insert user
func (api *ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("Username or password was no provided but is required.")
	}

	// check if user exists first
	exists, err := api.dbStore.DoesUserExist(event.Username)

	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("User exists")
	}

	// user does not exist, so add user
	err = api.dbStore.InsertUser(types.RegisterUser{
		Username: event.Username,
		Password: event.Password,
	})

	if err != nil {
		return nil
	}

	return nil
}
