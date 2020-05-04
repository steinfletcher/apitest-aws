package recorder

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/steinfletcher/apitest"
)

const source = "SNS"

func NewSNS(cli snsiface.SNSAPI, recorder *apitest.Recorder) snsiface.SNSAPI {
	return &snsRecorder{
		SNSAPI:   cli,
		recorder: recorder,
	}
}

type snsRecorder struct {
	snsiface.SNSAPI
	recorder *apitest.Recorder
}

func (r snsRecorder) Publish(input *sns.PublishInput) (*sns.PublishOutput, error) {
	recordInput(r.recorder, source, "PublishInput", input.String())

	output, err := r.SNSAPI.Publish(input)

	var body string
	if output != nil {
		body = output.String()
	}

	recordOutput(r.recorder, source, "PublishOutput", body, nil)

	return output, err
}
