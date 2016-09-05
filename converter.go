package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func converter(messages chan *sarama.ConsumerMessage) {
	log.Println("Converter ready")
	for true {
		msg := <-messages
		fmt.Println(string(msg.Value))
	}
}
