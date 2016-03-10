package main

import (
	"fmt"
	"log"

	"gopkg.in/Shopify/sarama.v1"
)

func converter(messages chan *sarama.ConsumerMessage) {
	log.Println("Converter ready")
	for true {
		msg := <-messages
		fmt.Println(string(msg.Value))
	}
}
