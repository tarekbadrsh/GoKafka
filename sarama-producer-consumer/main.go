package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

var topic = os.Getenv("topic")
var brokers = []string{os.Getenv("kafkaURL")}

func main() {
	fmt.Println(">>> topic ", topic)
	fmt.Println(">>> brokers ", brokers)
	producer, err := newProducer(brokers)
	if err != nil {
		fmt.Println("Could not create producer: ", err)
	}

	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		fmt.Println("Could not create consumer: ", err)
	}

	if err != nil {
		return
	}

	subscribe(topic, consumer)
	i := 0
	for {
		i++
		time.Sleep(time.Second)

		msg := prepareMessage(topic, fmt.Sprintf("MSG-%d ", i))
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Printf(" %s error occured.", err.Error())
		} else {
			fmt.Printf(" Message was saved to partion: %d. Message offset is: %d.\n", partition, offset)
		}
	}
}
