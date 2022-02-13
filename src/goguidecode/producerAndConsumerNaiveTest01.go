package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
// 创建一把互斥锁
var lock2 = sync.Mutex{}

// 定义缓冲区
var sce []int = make([]int, 10)

// 定义生产者
func producer(){
	// 加锁, 注意是lock就是我们的锁, 全局公用一把锁
	lock2.Lock()
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++{
		num := rand.Intn(100)
		sce[i] = num
		fmt.Println("生产者生产了: ", num)
		time.Sleep(time.Millisecond * 500)
	}
	// 解锁
	lock2.Unlock()
}
// 定义消费者
func consumer()  {
	// 加锁, 注意和生产者中用的是同一把锁
	// 如果生产者中已加过了, 则阻塞直到解锁后再重新加锁
	lock.Lock()
	for i:=0;i<10;i++{
		num := sce[i]
		fmt.Println("---消费者消费了", num)
	}
	lock.Unlock()
}

func main() {
	go producer()
	go consumer()
	for{
		;
	}
}