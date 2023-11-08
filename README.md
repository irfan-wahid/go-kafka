# go-kafka

This is an example of implementing Kafka producer and consumer using Golang.

## How to run this app
First, you must run zookeeper and kafka, you can download and see the documentation how to run kafka in this link

https://kafka.apache.org/

Second, open kafka and create a topic. you can see how to create a topic from that link too.

Third, open and run go-producers folder using command

    go run main.go

If you want to change the data that will be sent to the topic or change the topic name, you can change on file go-producers/kafka/kafka-config.go

Last, If you want to check whether the data has entered the topic or not, you can run go-consumer file using command

    go run main.go

You can change topic name on file go-consumer/kafka/kafka-config.go