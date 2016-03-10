package main

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/Shopify/sarama.v1"
)

func main() {
	log.SetOutput(os.Stderr)
	// Input kafka configuration
	var (
		inputKafka     = getConfig("INPUT_KAFKA", "localhost")  // The DNS name for input Kafka broker service
		inputKafkaPort = getConfig("INPUT_KAFKA_PORT", "9092")  // Port to connect to input Kafka peers
		topic          = getConfig("TOPIC", "create-agreement") // `create-agreement`. The topic to consume
		partitions     = getConfig("PARTITION", "all")          // `all`. The partitions to consume, can be 'all' or comma-separated numbers
		offset         = getConfig("OFFSET", "newest")          // `newest`. The offset to start with. Can be `oldest`, `newest`
		bufferSize     = getConfig("BUFFER_SIZE", "256")        // `256`. The buffer size of the message channel. Recommended `256`
	)
	// Common configuration
	var (
		verbose = getConfig("VERBOSE", "false") // `false` . `true` if you want to turn on sarama logging
	)

	messages := initializeChannels(bufferSize)
	go converter(messages)
	consumer(inputKafka, inputKafkaPort, topic, partitions, offset, messages, verbose == "true")
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
