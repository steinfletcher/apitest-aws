package recorder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/steinfletcher/apitest"
)

const outputLimitBytes = 1024

func NewS3(cli s3iface.S3API, recorder *apitest.Recorder) s3iface.S3API {
	return &s3Recorder{
		S3API:    cli,
		recorder: recorder,
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

func (r s3Recorder) ListObjects(input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	r.recordInput("ListObjectsInput", input.String())

	output, err := r.S3API.ListObjects(input)

	var body string
	if output != nil {
		body = output.String()
	}

	r.recordOutput("ListObjectsOutput", body, nil)

	return output, err
}

func (r s3Recorder) ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	r.recordInput("ListObjectsV2Input", input.String())

	output, err := r.S3API.ListObjectsV2(input)

	var body string
	if output != nil {
		body = output.String()
	}

	r.recordOutput("ListObjectsV2Output", body, nil)

	return output, err
}

func (r s3Recorder) HeadObject(input *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	r.recordInput("HeadObjectInput", input.String())

	output, err := r.S3API.HeadObject(input)

	var body string
	if output != nil {
		body = output.String()
	}

	r.recordOutput("HeadObjectOutput", body, nil)

	return output, err
}

func (r s3Recorder) CreateMultipartUpload(input *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	r.recordInput("CreateMultipartUploadInput", input.String())

	output, err := r.S3API.CreateMultipartUpload(input)

	var body string
	if output != nil {
		body = output.String()
	}

	r.recordOutput("CreateMultipartUploadOutput", body, err)

	return output, err
}

func (r s3Recorder) CompleteMultipartUpload(input *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	r.recordInput("CompleteMultipartUploadInput", input.String())

	output, err := r.S3API.CompleteMultipartUpload(input)

	var body string
	if output != nil {
		body = output.String()
	}

	r.recordOutput("CompleteMultipartUploadOutput", body, err)

	return output, err
}

func (r s3Recorder) UploadPartRequest(input *s3.UploadPartInput) (*request.Request, *s3.UploadPartOutput) {
	r.recordInput("UploadPartInput", input.String())

	req, output := r.S3API.UploadPartRequest(input)

	var body string
	if req != nil {
		body = output.String()
	}

	r.recordOutput("UploadPartOutput", body, nil)

	return req, output
}

func (r s3Recorder) ListParts(input *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	r.recordInput("ListPartsInput", input.String())

	output, err := r.S3API.ListParts(input)

	var body string
	if output != nil {
		body = output.String()
	}

	r.recordOutput("ListPartsOutput", body, nil)

	return output, err
}

func (r s3Recorder) AbortMultipartUpload(input *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	r.recordInput("AbortMultipartUploadInput", input.String())

	output, err := r.S3API.AbortMultipartUpload(input)

	var body string
	if output != nil {
		body = output.String()
	}

	r.recordOutput("AbortMultipartUploadOutput", body, nil)

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
