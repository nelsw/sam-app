package client

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"log"
	"os"
	"sam-app/pkg/client/faas"
)

var l *lambda.Lambda

func init() {
	if sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}); err != nil {
		log.Fatalf("Failed to connect to AWS: %s", err.Error())
	} else {
		l = lambda.New(sess)
	}
}

func Invoke(i faas.Functional) ([]byte, error) {
	payload, _ := json.Marshal(&i)
	if out, err := l.Invoke(&lambda.InvokeInput{
		FunctionName: aws.String(i.Function()),
		Payload:      payload,
	}); err != nil {
		return nil, err
	} else {
		return out.Payload, nil
	}
}

func InvokeRaw(b []byte, s string) ([]byte, error) {
	if out, err := l.Invoke(&lambda.InvokeInput{
		FunctionName: aws.String(s),
		Payload:      b,
	}); err != nil {
		return nil, err
	} else {
		return out.Payload, nil
	}
}

func Call(i interface{}, s string) ([]byte, error) {
	b, _ := json.Marshal(&i)
	if out, err := l.Invoke(&lambda.InvokeInput{
		FunctionName: aws.String(s),
		Payload:      b,
	}); err != nil {
		return nil, err
	} else {
		_ = json.Unmarshal(out.Payload, &err)
		return out.Payload, err
	}
}
