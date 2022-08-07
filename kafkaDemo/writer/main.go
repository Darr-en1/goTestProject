package main

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func sendMessage() {
	// 使用kafka.Writer kafka.NewWriter 废弃，不推荐使用
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "topic-A",
		Balancer: &kafka.LeastBytes{}, // 消息分区平衡器，用于将消息分配到多个分区的算法
		// Compression 参数指定压缩解压的解码器

		// 通过 Transport 设置 TLS
		//Transport: &kafka.Transport{
		//	TLS: &tls.Config{},
		//},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func sendMessageAutoTopicCreation() {
	// 如果缺少topic，则创建
	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "topic-A",
		AllowAutoTopicCreation: true,
	}

	messages := []kafka.Message{
		{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	}

	var err error
	const retries = 3
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for i := 0; i < retries; i++ {
		// attempt to create topic prior to publishing the message
		err = w.WriteMessages(ctx, messages...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			log.Fatalf("unexpected error %v", err)
		}
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

// 这两种模式是互斥的，如果你设置了Writer.Topic，你不能Message.Topic在你正在写的消息上明确定义。当您没有为作者定义主题时，情况正好相反。如果检测到这种歧义，Writer它将返回错误。
func sendMessageWithTopic() {
	w := &kafka.Writer{
		Addr: kafka.TCP("localhost:9092"),
		// 如果此处未定义Topic，则每个消息都必须定义它
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		// 每个 Message 都定义Topic，否则返回错误。
		kafka.Message{
			Topic: "topic-A",
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Topic: "topic-B",
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Topic: "topic-C",
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func main() {
	sendMessage()
	//sendMessageAutoTopicCreation()
	//sendMessageWithTopic()
}
