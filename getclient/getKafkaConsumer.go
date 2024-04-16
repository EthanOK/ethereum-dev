package getclient

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func GetKafkaConsumer(consumer_group, kafka_key, kafka_secret string) *kafka.Consumer {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "sync-streaming.chainbase.online:9093",
		"security.protocol": "SASL_PLAINTEXT",
		"sasl.mechanisms":   "SCRAM-SHA-256",
		"group.id":          consumer_group,
		"sasl.username":     kafka_key,
		"sasl.password":     kafka_secret,
		"auto.offset.reset": "earliest",
		"socket.timeout.ms": 10000,
	})

	if err != nil {
		panic(err)
	}
	return c

}
