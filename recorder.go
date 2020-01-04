package dynamodbrecorder

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/steinfletcher/apitest"
)

func New(cli dynamodbiface.DynamoDBAPI, recorder *apitest.Recorder) dynamodbiface.DynamoDBAPI {
	return &dynamoRecorder{
		DynamoDBAPI: cli,
		recorder:    recorder,
	}
}

type dynamoRecorder struct {
	dynamodbiface.DynamoDBAPI
	recorder *apitest.Recorder
}

func (a dynamoRecorder) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	a.recordInput("QueryInput", input.String())

	output, err := a.DynamoDBAPI.Query(input)

	var body string
	if output != nil {
		body = output.String()
	}

	a.recordOutput("QueryOutput", body, err)

	return output, err
}

func (a dynamoRecorder) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	a.recordInput("UpdateItemInput", input.String())

	output, err := a.DynamoDBAPI.UpdateItem(input)

	var body string
	if output != nil {
		body = output.String()
	}

	a.recordOutput("UpdateItemOutput", body, err)

	return output, err
}

func (a dynamoRecorder) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	a.recordInput("PutItemInput", input.String())

	output, err := a.DynamoDBAPI.PutItem(input)

	var body string
	if output != nil {
		body = output.String()
	}

	a.recordOutput("PutItemOutput", body, err)

	return output, err
}

func (a dynamoRecorder) recordInput(operation, body string) {
	a.recorder.AddMessageRequest(apitest.MessageRequest{
		Source:    apitest.SystemUnderTestDefaultName,
		Target:    "DynamoDB",
		Header:    operation,
		Body:      body,
		Timestamp: time.Now(),
	})
}

func (a dynamoRecorder) recordOutput(operation, body string, err error) {
	if err != nil {
		a.recorder.AddMessageResponse(apitest.MessageResponse{
			Source:    "DynamoDB",
			Target:    apitest.SystemUnderTestDefaultName,
			Header:    "Error",
			Body:      fmt.Sprintf("Error: %s", err.Error()),
			Timestamp: time.Now(),
		})
	} else {
		a.recorder.AddMessageResponse(apitest.MessageResponse{
			Source:    "DynamoDB",
			Target:    apitest.SystemUnderTestDefaultName,
			Header:    operation,
			Body:      body,
			Timestamp: time.Now(),
		})
	}
}
