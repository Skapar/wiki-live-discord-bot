package messaging

import (
	"github.com/IBM/sarama"
	"github.com/Skapar/wiki-live-discord-bot/internal/model"
)

type KafkaProducer struct {
    producer sarama.SyncProducer
}

func NewKafkaProducer(brokers string) (*KafkaProducer, error) {
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
    config.Producer.Return.Successes = true

    producer, err := sarama.NewSyncProducer([]string{brokers}, config)
    if err != nil {
        return nil, err
    }

    return &KafkaProducer{producer: producer}, nil
}

func (kp *KafkaProducer) SendMessage(message *model.Message) error {
    msg := &sarama.ProducerMessage{
        Topic: message.Topic,
        Value: sarama.StringEncoder(message.Value),
    }
    _, _, err := kp.producer.SendMessage(msg)
    return err
}

func (kp *KafkaProducer) Close() error {
    return kp.producer.Close()
}