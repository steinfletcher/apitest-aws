package examples

import (
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/steinfletcher/apitest"

	"github.com/steinfletcher/apitest-aws/recorder"
)

func GetUser(t *testing.T) {
	test, db := apiTest()

	handler := http.NewServeMux()
	handler.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, _ = db.PutItem(&dynamodb.PutItemInput{
			TableName: aws.String("my_table"),
			Item:      nil, // Add item here
		})
		w.WriteHeader(http.StatusCreated)
	})

	test.Handler(handler).
		Get("/hello").
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func apiTest() (*apitest.APITest, dynamodbiface.DynamoDBAPI) {
	rec := apitest.NewTestRecorder()
	db := recordingDB(rec)

	return apitest.New().
		Recorder(rec).
		Report(apitest.SequenceDiagram()), db
}

func recordingDB(rec *apitest.Recorder) dynamodbiface.DynamoDBAPI {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("localhost"),
		Endpoint: aws.String("http://localhost:8000"),
	}))
	db := dynamodb.New(sess)
	return recorder.NewDynamoDB(db, rec)
}
