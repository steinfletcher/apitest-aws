package recorder

import (
	"fmt"
	"time"

	"github.com/steinfletcher/apitest"
)

func recordInput(recorder *apitest.Recorder, source, operation, body string) {
	recorder.AddMessageRequest(apitest.MessageRequest{
		Source:    apitest.SystemUnderTestDefaultName,
		Target:    source,
		Header:    operation,
		Body:      body,
		Timestamp: time.Now(),
	})
}

func recordOutput(recorder *apitest.Recorder, source, operation, body string, err error) {
	if err != nil {
		recorder.AddMessageResponse(apitest.MessageResponse{
			Source:    source,
			Target:    apitest.SystemUnderTestDefaultName,
			Header:    "Error",
			Body:      fmt.Sprintf("Error: %s", err.Error()),
			Timestamp: time.Now(),
		})
	} else {
		recorder.AddMessageResponse(apitest.MessageResponse{
			Source:    source,
			Target:    apitest.SystemUnderTestDefaultName,
			Header:    operation,
			Body:      body,
			Timestamp: time.Now(),
		})
	}
}
