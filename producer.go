package main

import (        
        "github.com/segmentio/kafka-go"
        "context"
)

func main() {
        topic := "my-topic"
        
        w := kafka.NewWriter(kafka.WriterConfig{
                Brokers: []string{"localhost:9092"},
                Topic:   topic,
                Balancer: &kafka.LeastBytes{},
        })

        w.WriteMessages(context.Background(),
                kafka.Message{
                        Value: []byte("Hello World!"),
                },
                kafka.Message{
                        Key:   []byte("Key-A"),
                        Value: []byte("One!"),
                },
                kafka.Message{
                        Key:   []byte("Key-B"),
                        Value: []byte("Two!"),
                },
        )

        w.Close()
}
