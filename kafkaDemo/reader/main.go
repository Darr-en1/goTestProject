package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func readMessage() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "topic-A",
		//GroupID: "reader-003", // 设置GroupID ,使用 ConsumerGroup 方式消费
		Partition: 0,                // 从指定分区获取数据
		MinBytes:  1,                // 10KB
		MaxBytes:  10e6,             // 10MB
		MaxWait:   20 * time.Second, // 为达到MinBytes的读取时间间隔
	})
	r.SetOffset(0) // GroupID 未设置可生效

	// r.SetOffsetAt // 通过SetOffsetAt可以设置从什么时候产生的消息开始消费

	for {
		m, err := r.ReadMessage(context.Background()) // consumerGroup 会自动提交，一条一条同步提交offset,效率低下，如果后续业务未执行（如断电了）则会丢失
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func readMessageWithGroupManuallyCommit() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer-group-id-001", //  设置GroupID ,使用 ConsumerGroup 方式消费
		Topic:    "topic-A",
		MinBytes: 1,    // 10KB
		MaxBytes: 10e6, // 10MB
		//CommitInterval: time.Second, // 设置提交间隔， CommitMessages 会变成异步提交
	})
	background := context.Background()

	for {
		m, err := r.FetchMessage(background) // FetchMessage方法读取 consumerGroup 不会提交offset
		if err != nil {
			break
		}
		// 处理业务逻辑后，决定是否要提交offset,保证的消息的不被丢失，但是会存在重复消费（重复数量为一），需要在业务中去重
		// 一条一条提交 效率还是很差，可以采取多条消费后在提交最后一条的offset,最好的方式是设置 CommitInterval 参数 采取异步提交方式，但这两种都会操作更多的重复消费
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		if err = r.CommitMessages(background, m); err != nil {
			break
		}

	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func readMessageWithGroupAutoCommit() {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer-group-id", // 设置GroupID ,使用 ConsumerGroup 方式消费
		Topic:    "topic-A",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background()) // ReadMessage方法读取 consumerGroup 会自动提交，一条一条提交，如果后续业务未执行（如断电了）则会丢失
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func main() {
	//readMessage()
	//readMessageWithGroupAutoCommit()
	readMessageWithGroupManuallyCommit()
}
