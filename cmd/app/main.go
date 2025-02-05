package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	fmt.Println("WIKI LIVE BOT is starting...")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Ошибка подключения к Redis: %v\n", err)
	} else {
		fmt.Println("Успешное подключение к Redis!")
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		log.Fatalf("Ошибка подключения к Kafka: %v\n", err)
	} else {
		fmt.Println("Успешное подключение к Kafka!")
	}

	message := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("Hello Kafka!"),
	}
	_, _, err = producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Ошибка отправки сообщения в Kafka: %v\n", err)
	} else {
		fmt.Println("Сообщение успешно отправлено в Kafka!")
	}

	defer producer.Close()

	time.Sleep(5 * time.Second)
}
