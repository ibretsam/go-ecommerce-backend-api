package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL = "localhost:9092"
	topic    = "VipUserTopic"
)

// For Producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

// For Consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// Buy or Sell Stock
func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg

	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)

	// Create message
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	// Write message to Kafka
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to send message to Kafka",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success send message to Kafka",
	})
}

func RegisterConsumerATC(id int) {
	kafkaGroupId := "consumer-group-"
	reader := getKafkaReader(kafkaURL, topic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer %d waiting at ATC:: \n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("Consumer %d, with topic %v, partition %v, offset %v at time %d %s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, topic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)

	// Register 2 user for buying stock in ATC
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)

	r.Run(":8999")
}
