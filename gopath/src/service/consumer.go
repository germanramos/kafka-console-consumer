package main

import (
	"errors"
	"log"
	"net"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"gopkg.in/bsm/sarama-cluster.v2"
)

func consumer(kafkaService string,
	kafkaPort string,
	topic string,
	groupID string,
	offset string,
	messages chan *sarama.ConsumerMessage,
	verbose bool) {
	var (
		err        error
		brokerList []string
		consumer   *cluster.Consumer
	)

	if verbose == true {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	var initialOffset int64
	switch offset {
	case "oldest":
		initialOffset = sarama.OffsetOldest
	case "newest":
		initialOffset = sarama.OffsetNewest
	default:
		log.Fatalln("Offset should be `oldest` or `newest`")
	}

	// Get Kafka peers
	for err != nil || brokerList == nil {
		brokerList, err = net.LookupHost(kafkaService)
		if err != nil {
			log.Printf("Failed to resolve %s: %s\n", kafkaService, err)
			time.Sleep(time.Second * 3)
		}
	}
	for i, e := range brokerList {
		brokerList[i] = e + ":" + kafkaPort
	}
	log.Println("Input Kafka Broker List:", brokerList)

	config := cluster.NewConfig()
	config.Consumer.Offsets.Initial = initialOffset
	err = errors.New("Init")
	for err != nil {
		consumer, err = cluster.NewConsumer(brokerList, groupID, []string{topic}, config)
		if err != nil {
			log.Println("Failed to start Sarama consumer:", err)
			time.Sleep(time.Second * 3)
		} else {
			log.Println("Consumer ready")
		}
	}

	for msg := range consumer.Messages() {
		messages <- msg
		consumer.MarkOffset(msg, "")
	}

}
