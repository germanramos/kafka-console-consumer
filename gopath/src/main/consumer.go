package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"

	"gopkg.in/Shopify/sarama.v1"
)

func consumer(inputKafka string,
	inputKafkaPort string,
	topic string,
	partitions string,
	offset string,
	messages chan *sarama.ConsumerMessage,
	verbose bool) {
	var (
		err        error
		brokerList []string
		c          sarama.Consumer
		closing    = make(chan struct{})
		wg         sync.WaitGroup
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
		brokerList, err = net.LookupHost(inputKafka)
		if err != nil {
			log.Printf("Failed to resolve %s: %s\n", inputKafka, err)
			time.Sleep(time.Second * 3)
		}
	}
	for i, e := range brokerList {
		brokerList[i] = e + ":" + inputKafkaPort
	}
	log.Println("Input Kafka Broker List:", brokerList)

	// Create consumer
	err = nil
	for c == nil {
		c, err = sarama.NewConsumer(brokerList, nil)
		if err != nil {
			log.Println("Failed to start Sarama consumer:", err)
			time.Sleep(time.Second * 3)
		}
	}

	partitionList, err := getPartitions(c, topic, partitions)
	if err != nil {
		log.Fatalln("Failed to get the list of partitions: %s", err)
	}

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Kill, os.Interrupt)
		<-signals
		log.Println("Initiating shutdown of consumer...")
		close(closing)
	}()

	// Create a go routine to consume every partition
	for _, partition := range partitionList {
		pc, err := c.ConsumePartition(topic, partition, initialOffset)
		if err != nil {
			log.Fatalln("Failed to start consumer for partition %d: %s", partition, err)
		}

		go func(pc sarama.PartitionConsumer) {
			<-closing
			pc.AsyncClose()
		}(pc)

		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			log.Println("Consumer ready")
			for message := range pc.Messages() {
				messages <- message
			}
		}(pc)
	}

	// Wait for consumming ends
	wg.Wait()
	log.Println("Done consuming topic", topic)
	close(messages)

	if err := c.Close(); err != nil {
		log.Println("Failed to close consumer: ", err)
	}
}

func getPartitions(c sarama.Consumer, topic string, partitions string) ([]int32, error) {
	if partitions == "all" {
		return c.Partitions(topic)
	}

	tmp := strings.Split(partitions, ",")
	var pList []int32
	for i := range tmp {
		val, err := strconv.ParseInt(tmp[i], 10, 32)
		if err != nil {
			return nil, err
		}
		pList = append(pList, int32(val))
	}

	return pList, nil
}
