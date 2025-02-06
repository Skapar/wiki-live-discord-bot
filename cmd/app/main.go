package main

import (
	"fmt"
	"time"

	"github.com/Skapar/wiki-live-discord-bot/internal/config"
	"github.com/Skapar/wiki-live-discord-bot/internal/delivery/gateway/messaging"
	"github.com/Skapar/wiki-live-discord-bot/internal/model"
	"github.com/Skapar/wiki-live-discord-bot/internal/repository"
	"github.com/Skapar/wiki-live-discord-bot/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
    cfg := config.New()
    cfg.Init()

    log := logger.New()

    fmt.Println("WIKI LIVE BOT is starting...")

    _, err := repository.NewRedisClient(cfg.RedisAddr)
    if err != nil {
        log.Fatalf("Ошибка подключения к Redis: %v\n", err)
    } else {
        log.Info("Успешное подключение к Redis!")
    }

    kafkaProducer, err := messaging.NewKafkaProducer(cfg.KafkaAddr)
    if err != nil {
        log.Fatalf("Ошибка подключения к Kafka: %v\n", err)
    } else {
        log.Info("Успешное подключение к Kafka!")
    }

    message := &model.Message{
        Topic: "test-topic",
        Value: "Hello Kafka!",
    }
    err = kafkaProducer.SendMessage(message)
    if err != nil {
        log.Fatalf("Ошибка отправки сообщения в Kafka: %v\n", err)
    } else {
        log.Info("Сообщение успешно отправлено в Kafka!")
    }

    defer kafkaProducer.Close()

    time.Sleep(5 * time.Second)
}