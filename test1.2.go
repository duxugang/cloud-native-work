package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
基于 Channel 编写一个简单的单线程生产者消费者模型：
队列：
队列长度 10，队列元素类型为 int
生产者：
每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/
func main() {
	queue := make(chan int, 10)
	go consumer(queue)
	product(queue)
}

func consumer(queue <-chan int) {
	timer := time.NewTimer(time.Second)
	for range timer.C {
		fmt.Printf("消费数据：[%v]\n", <-queue)
		timer.Reset(time.Second)
	}
}

func product(queue chan<- int) {
	ticker := time.NewTicker(time.Second).C
	for range ticker {
		queue <- rand.Intn(100)
	}
}
