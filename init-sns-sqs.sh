#!/bin/bash

set -e

# Iniciar o LocalStack
docker-compose up -d
# Esperar o LocalStack iniciar
sleep 5

# Criar tópico SNS
aws --endpoint-url=http://localhost:4566 sns create-topic --name MyTopic

# Criar fila SQS
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name MyQueue

# Pegar os ARNs do tópico SNS e da fila SQS
TOPIC_ARN=$(aws --endpoint-url=http://localhost:4566 sns list-topics --query 'Topics[0].TopicArn' --output text)
QUEUE_URL=$(aws --endpoint-url=http://localhost:4566 sqs get-queue-url --queue-name MyQueue --query 'QueueUrl' --output text)
QUEUE_ARN=$(aws --endpoint-url=http://localhost:4566 sqs get-queue-attributes --queue-url $QUEUE_URL --attribute-name QueueArn --query 'Attributes.QueueArn' --output text)

# Assinar a fila SQS ao tópico SNS
aws --endpoint-url=http://localhost:4566 sns subscribe --topic-arn $TOPIC_ARN --protocol sqs --notification-endpoint $QUEUE_ARN

echo "SNS Topic ARN: $TOPIC_ARN"
echo "SQS Queue URL: $QUEUE_URL"
echo "SQS Queue ARN: $QUEUE_ARN"
