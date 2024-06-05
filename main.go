package main

import (
	snsclient "github/LordRadamanthys/teste_localstack/client/sns_client"
	sqsclient "github/LordRadamanthys/teste_localstack/client/sqs_client"
	"sync"
	"time"
)

const (
	SNS_ARN   = "arn:aws:sns:us-east-1:000000000000:test-sns"
	LOCALHOST = "http://localhost:4566"
	SQS_URL   = "http://localhost:4566/000000000000/sqs-test"
	REGION    = "us-east-1"
)

func main() {

	go func() {
		for {
			snsclient.SendMessageSNS(REGION, LOCALHOST, SNS_ARN)
			time.Sleep(10 * time.Second)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go sqsclient.ReadMessageSQS(&wg, REGION, LOCALHOST, SQS_URL)
	wg.Wait()

}
