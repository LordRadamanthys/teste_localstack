- docker command using localstack: docker run --rm -d -p 4566:4566 --name localstack-test localstack/localstack

- create sns: aws sns --endpoint-url http://localhost:4566 create-topic --name test-sns 

- post a message: aws --endpoint-url=http://localhost:4566 sns publish --topic-arn arn:aws:sns:us-east-1:000000000000:test-sns --message "Jorge"

- subscribe sqs to sns: aws --endpoint-url=http://localhost:4566 sns subscribe --topic-arn $TOPIC_ARN --protocol sqs --notification-endpoint $QUEUE_ARN 

- create queue: aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name sqs-test
 

 run project go run main.go