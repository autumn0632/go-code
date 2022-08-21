package main

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func producer(value string) {
	w := &kafka.Writer{}
	w.Addr = kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094")
	// w.Topic = "topic-A"
	// 分区分配策略，默认轮询
	w.Balancer = &kafka.LeastBytes{}
	// 自动创建topic
	w.AllowAutoTopicCreation = true

	defer w.Close()

	w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: "topic",
			Key:   []byte("A"),
			Value: []byte(value)},
		// kafka.Message{Key: []byte("B"), Value: []byte("world")},
	)

}
