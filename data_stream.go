package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DataStreamStackProps struct {
	awscdk.StackProps
}

func NewDataStreamStack(scope constructs.Construct, id string, props *DataStreamStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// create new table on dynamodb
	table := awsdynamodb.NewTable(stack, jsii.String("user"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			// since we have username as our "primary key"
			Name: jsii.String("username"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName: jsii.String("userTable"), // matches the declaration inside database.go
	})

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("DataStreamQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	// a construct to add a lambda function
	// declare as a function to pass into the User Table's GrantReadWriteData
	myLambdaFn := awslambda.NewFunction(stack,
		// the name of the lambda function once deployed
		jsii.String("myLambdaFunction"), &awslambda.FunctionProps{
			Runtime: awslambda.Runtime_PROVIDED_AL2023(),
			Code:    awslambda.AssetCode_FromAsset(jsii.String("lambda/function.zip"), nil),
			Handler: jsii.String("main"),
		})

	// grants permission for our lambda function to read and write to the User Table
	table.GrantReadWriteData(myLambdaFn)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewDataStreamStack(app, "DataStreamStack", &DataStreamStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
