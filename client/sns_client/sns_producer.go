package snsclient

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendMessageSNS(region, path, arn string) {

	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(path),
	}))

	snsService := sns.New(sess)
	_, err := snsService.Publish(&sns.PublishInput{
		TopicArn: aws.String(arn),
		Message:  aws.String("Hello, World!"),
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Message sent!")
}
