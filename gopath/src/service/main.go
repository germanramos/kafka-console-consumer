package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Shopify/sarama"
)

func main() {
	log.SetOutput(os.Stderr)
	var (
		kafkaService = getConfig("KAFKA_SERVICE", "kafka")                // The DNS name for input Kafka broker service
		kafkaPort    = getConfig("KAFKA_PORT", "9092")                    // Port to connect to input Kafka peers
		topic        = getConfig("TOPIC", "create-agreement")             // `create-agreement`. The topic to consume
		groupID      = getConfig("KAFKA_GROUP", "kafka-console-consumer") // `all`. The partitions to consume, can be 'all' or comma-separated numbers
		offset       = getConfig("OFFSET", "newest")                      // `newest`. The offset to start with. Can be `oldest`, `newest`
		bufferSize   = getConfig("BUFFER_SIZE", "256")                    // `256`. The buffer size of the message channel. Recommended `256`
		verbose      = getConfig("VERBOSE", "false")                      // `false` . `true` if you want to turn on sarama logging
	)

	messages := initializeChannels(bufferSize)
	go converter(messages)
	consumer(kafkaService, kafkaPort, topic, groupID, offset, messages, verbose == "true")
}

func getConfig(key string, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		result = defaultValue
	}
	log.Printf("%s=%s\n", key, result)
	return result
}

func initializeChannels(bufferSize string) chan *sarama.ConsumerMessage {
	intputBufferSize, err := strconv.Atoi(bufferSize)
	if err != nil {
		log.Fatalln("Failed to parse input buffer size")
	}
	messages := make(chan *sarama.ConsumerMessage, intputBufferSize)
	return messages
}
