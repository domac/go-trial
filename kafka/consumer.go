package main

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	master, err := sarama.NewConsumer([]string{"vm-kafka:9092"}, config)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	consumer, err := master.ConsumePartition("test", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln(err)
	}

	var (
		wg       sync.WaitGroup
		msgCount int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for message := range consumer.Messages() {
			log.Printf("Consumed message : %s", string(message.Value))
			msgCount++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range consumer.Errors() {
			log.Println(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
	consumer.AsyncClose()
	wg.Wait()

	log.Println("Processed", msgCount, "messages.")
}
