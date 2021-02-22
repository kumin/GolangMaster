package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
)

var (
	exps = []string{uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String()}
)

type UserLogging struct {
	UserId       int
	EventId      int
	ExperimentId string
	VariantId    int
}

func PushKafkaData(ctx context.Context) {
	c := sarama.NewConfig()
	c.Producer.Compression = sarama.CompressionSnappy
	c.Producer.Partitioner = sarama.NewHashPartitioner
	c.Producer.Return.Successes = true
	c.Producer.MaxMessageBytes = 5242880 //5M
	producer, err := sarama.NewSyncProducer(
		[]string{"localhost:9092"},
		c,
	)
	if err != nil {
		log.Printf("Create Producer Error %v", err)
		panic(err)
	}
	defer producer.Close()

	for {
		logging := &UserLogging{
			UserId:       rand.Intn(1000),
			EventId:      rand.Intn(10),
			ExperimentId: exps[rand.Intn(4)],
			VariantId:    rand.Intn(4),
		}
		data, _ := json.Marshal(logging)
		fmt.Printf("Message: %s\n", string(data))
		kafkaMessage := &sarama.ProducerMessage{
			Topic: "user_testing",
			Key:   sarama.StringEncoder(strconv.Itoa(logging.UserId)),
			Value: sarama.StringEncoder(string(data)),
		}
		producer.SendMessage(kafkaMessage)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	ctx := context.Background()
	PushKafkaData(ctx)
}
