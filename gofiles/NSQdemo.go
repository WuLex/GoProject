package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/nsqio/go-nsq"
)

func main() {
	// 创建一个NSQ生产者
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个NSQ消费者
	consumer, err := nsq.NewConsumer("topic", "channel", nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	// 设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Received message: %s", message.Body)
		return nil
	}))

	// 连接到NSQD
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatal(err)
	}

	// 等待程序退出信号
	wg := &sync.WaitGroup{}
	wg.Add(1)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		producer.Stop()
		consumer.Stop()
		wg.Done()
	}()

	// 发送一条消息到NSQD
	err = producer.Publish("topic", []byte("Hello NSQ!"))
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
