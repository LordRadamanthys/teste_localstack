execute sh script ./init-sns-sqs.sh

the script will start a docker with localstack image using docker-compose
after that the script will create sns topic and sqs queue and make de subscription the sqs on sns
