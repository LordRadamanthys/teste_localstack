package sqsclient

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ReadMessageSQS(wg *sync.WaitGroup, region, path, sqsUrl string) {
	defer wg.Done()
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(path),
	}))
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(sqsUrl),
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(10),
	}
	sqsService := sqs.New(sess)
	for {

		result, err := sqsService.ReceiveMessage(params)
		if err != nil {
			fmt.Println(err.Error())
		}
		if len(result.Messages) == 0 {
			fmt.Println("No messages received")
			continue
		}
		for _, message := range result.Messages {
			fmt.Printf("Message ID: %s, Body: %s\n", *message.MessageId, *message.Body)

			DeleteSqsMessage(sqsService, message, sqsUrl)
		}
		fmt.Println("Message deleted successfully")
	}

}

func DeleteSqsMessage(sqsService *sqs.SQS, message *sqs.Message, sqsUrl string) {
	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(sqsUrl),
		ReceiptHandle: message.ReceiptHandle,
	}

	_, err := sqsService.DeleteMessage(deleteParams)
	if err != nil {
		fmt.Println("Delete Error", err)
		return
	}
}
