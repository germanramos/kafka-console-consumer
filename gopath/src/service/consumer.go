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

func getInitialOffset(offset string) int64 {
	var initialOffset int64
	switch offset {
	case "oldest":
		initialOffset = sarama.OffsetOldest
	case "newest":
		initialOffset = sarama.OffsetNewest
	default:
		log.Fatalln("Offset should be `oldest` or `newest`")
	}
	return initialOffset
}

func getKafkaPeers(kafkaService string, kafkaPort string) []string {
	var err error
	var brokerList []string
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
	return brokerList
}

func waitForTopic(brokerList []string, topic string) {
	for {
		saramaConfig := sarama.NewConfig()
		client, _ := sarama.NewClient(brokerList, saramaConfig)
		topics, _ := client.Topics()
		for _, topicElement := range topics {
			if topic == topicElement {
				log.Printf("Topic %s found\n", topic)
				return
			}
		}
		log.Printf("Failed to find topic %s\n", topic)
		time.Sleep(time.Second * 3)
	}
}

func consumer(kafkaService string,
	kafkaPort string,
	topic string,
	groupID string,
	offset string,
	messages chan *sarama.ConsumerMessage,
	verbose bool) {

	var consumer *cluster.Consumer

	if verbose == true {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	initialOffset := getInitialOffset(offset)
	brokerList := getKafkaPeers(kafkaService, kafkaPort)
	waitForTopic(brokerList, topic)

	// Create consumer with sarama-cluster
	clusterConfig := cluster.NewConfig()
	clusterConfig.Consumer.Offsets.Initial = initialOffset
	err := errors.New("Init")
	for err != nil {
		consumer, err = cluster.NewConsumer(brokerList, groupID, []string{topic}, clusterConfig)
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
