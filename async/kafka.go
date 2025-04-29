package async

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
	"gocode.ethan/ethereum-dev/getclient"
)

func StartKafka() {
	// https://console.chainbase.com/sync/kafka

	if err := godotenv.Load(); err != nil {
		fmt.Println("load .env failed")
	}

	consumer_group := os.Getenv("CONSUMER_GROUP")
	kafka_key := os.Getenv("KAFKA_KEY")
	kafka_secret := os.Getenv("KAFKA_SECRET")

	c := getclient.GetKafkaConsumer(consumer_group, kafka_key, kafka_secret)

	// []string{"ethereum_contracts", "ethereum.erc721.transfer"}
	c.SubscribeTopics([]string{"ethereum_contracts"}, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: \n", msg.TopicPartition)
			fmt.Println(string(msg.Value))

			// string to json
			// var data map[string]interface{}
			// json.Unmarshal(msg.Value, &data)
			// fmt.Println(data["type"], data["block_number"], ":", data["address"])
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
