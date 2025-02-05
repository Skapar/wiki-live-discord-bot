package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Skapar/wiki-live-discord-bot/internal/config"

	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
)

func main() {
	err := godotenv.Load()

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "trace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
	}
	zlogger := zap.New(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zap.DebugLevel),
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	log := zlogger.Sugar()

	cfg := config.New()
	cfg.Init()


	fmt.Println("WIKI LIVE BOT is starting...")
	
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
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

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	kafkaConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{cfg.KafkaAddr}, kafkaConfig)
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
