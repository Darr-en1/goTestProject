package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func consume() {
	topic := "my-topic-001"
	partition := 0 //  topic下的某一个partition 建立的连接。

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	// 设置 Read 调用的截止日期，timeout时间,批量读取时，没有达到最小量但超时了也会向下执行
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	//conn.SetReadDeadline(time.Time{}) // 不设置超时
	batch := conn.ReadBatch(10e3, 1e6) // 批量读取，读取量fetch 10KB min, 1MB max
	b := make([]byte, 10e3)            // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		partition := batch.Partition()
		offset := batch.Offset()
		fmt.Println(partition, offset)
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

func main() {
	consume()
}
