package database

import (
	"fmt"
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TABLE_NAME = "userTable"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() DynamoDBClient {
	// create a dynamo db session and pass it down
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return DynamoDBClient{
		databaseStore: db,
	}
}

// querying if user exists
func (u *DynamoDBClient) DoesUserExist(username string) (bool, error) {

	// use dynamodb lib to create the get object to acquire the item
	getItemKey := &dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	}

	result, err := u.databaseStore.GetItem(getItemKey)

	if err != nil {
		return true, err
	}

	if result.Item != nil {
		return false, nil
	}

	return true, nil
}

// inserting user
func (u *DynamoDBClient) InsertUser(user types.RegisterUser) error {
	// all the records we want to insert
	newUser := map[string]*dynamodb.AttributeValue{
		"username": {
			S: aws.String(user.Username),
		},
		"password": {
			S: aws.String(user.Password),
		},
	}

	// assemble the client of insertion
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item:      newUser,
	}

	_, err := u.databaseStore.PutItem(item)

	if err != nil {
		fmt.Println("Inserting item did not succeed due to error:", err)
		return err
	}

	return nil
}
