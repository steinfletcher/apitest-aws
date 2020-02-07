package recorder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/steinfletcher/apitest"
)

const outputLimitBytes = 1024

func NewS3(cli s3iface.S3API, recorder *apitest.Recorder) s3iface.S3API {
	return &s3Recorder{
		S3API: cli,
		recorder:    recorder,
	}
}

type s3Recorder struct {
	s3iface.S3API
	recorder *apitest.Recorder
}

func (r s3Recorder) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	r.recordInput("GetObjectInput", input.String())

	output, err := r.S3API.GetObject(input)

	var body string
	if output != nil {
		if output.Body != nil {
			data, err := ioutil.ReadAll(output.Body)
			if err != nil {
				panic(err)
			}

			output.Body = ioutil.NopCloser(bytes.NewBuffer(data))

			if len(data) < outputLimitBytes {
				body = string(data)
			} else {
				body = fmt.Sprintf("Content Length: %d", len(data))
			}
		}
	}

	r.recordOutput("GetObjectOutput", body, err)

	return output, err
}

func (r s3Recorder) recordInput(operation, body string) {
	r.recorder.AddMessageRequest(apitest.MessageRequest{
		Source:    apitest.SystemUnderTestDefaultName,
		Target:    "S3",
		Header:    operation,
		Body:      body,
		Timestamp: time.Now(),
	})
}

func (r s3Recorder) recordOutput(operation, body string, err error) {
	if err != nil {
		r.recorder.AddMessageResponse(apitest.MessageResponse{
			Source:    "S3",
			Target:    apitest.SystemUnderTestDefaultName,
			Header:    "Error",
			Body:      fmt.Sprintf("Error: %s", err.Error()),
			Timestamp: time.Now(),
		})
	} else {
		r.recorder.AddMessageResponse(apitest.MessageResponse{
			Source:    "S3",
			Target:    apitest.SystemUnderTestDefaultName,
			Header:    operation,
			Body:      body,
			Timestamp: time.Now(),
		})
	}
}