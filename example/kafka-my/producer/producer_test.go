package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/segmentio/kafka-go"
)

func TestProducer(t *testing.T) {

	w := &kafka.Writer{}
	w.Addr = kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094")
	w.Topic = "topic-A"
	// 分区分配策略，默认轮询
	w.Balancer = &kafka.LeastBytes{}
	w.Async = true
	w.AllowAutoTopicCreation = true

	err := w.WriteMessages(context.Background(),
		kafka.Message{Key: []byte("A"), Value: []byte("world")},
	)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkProducer(b *testing.B) {
	w := &kafka.Writer{}
	w.Addr = kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094")
	// w.Topic = "topic-A"
	// 分区分配策略，默认轮询
	// w.Balancer = &kafka.LeastBytes{}
	w.Async = true
	w.AllowAutoTopicCreation = true

	defer w.Close()
	for i := 0; i < b.N; i++ {
		w.WriteMessages(context.Background(),
			kafka.Message{
				Topic: fmt.Sprintf("topic-%d", i),
				Key:   []byte("B"),
				Value: []byte("world")},
		)
	}

}
