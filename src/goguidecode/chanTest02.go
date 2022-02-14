package main

import (
	"fmt"
	"math/rand"
	"time"
)
// 定义缓冲区
var myCh = make(chan int, 5)
var exitCh = make(chan bool, 1)

// 定义生产者
func producer2(){
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++{
		num := rand.Intn(100)
		fmt.Println("生产者生产了: ", num)
		// 往管道中写入数据
		myCh<-num
		time.Sleep(time.Millisecond * 100)
	}
	// 生产完毕之后关闭管道
	close(myCh)
	fmt.Println("生产者停止生产")
}
// 定义消费者
func consumer2()  {
	// 不断从管道中获取数据, 直到管道关闭位置
	for{
		if num, ok := <-myCh; !ok{
			break
		}else{
			fmt.Println("---消费者消费了", num)
		}
	}
	fmt.Println("消费者停止消费")
	exitCh<-true
}

func main() {
	go producer2()
	go consumer2()
	fmt.Println("exitCh之前代码")
	// 这个是在阻塞。等待结束。
	<-exitCh
	fmt.Println("exitCh之后代码")

	go producer3()
	go consumer3()
	<- exitCh2

	var myCh4 = make(chan int, 5)
	go producer4(myCh4)
	consumer4(myCh4)
}


// 定义无缓冲Channel
//var myCh = make(chan int, 0)
var myCh2 = make(chan int, 0)
var exitCh2 = make(chan bool, 1)

// 定义生产者
func producer3(){
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++{
		num := rand.Intn(100)
		fmt.Println("生产者生产了: ", num)
		// 往管道中写入数据
		myCh2<-num
		//time.Sleep(time.Millisecond * 500)
	}
	// 生产完毕之后关闭管道
	close(myCh2)
	fmt.Println("生产者停止生产")
}
// 定义消费者
func consumer3()  {
	// 不断从管道中获取数据, 直到管道关闭位置
	for{
		if num, ok := <-myCh2; !ok{
			break
		}else{
			fmt.Println("---消费者消费了", num)
		}
	}
	fmt.Println("消费者停止消费")
	exitCh2<-true
}

// 定义生产者
func producer4(myCh chan<- int){
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++{
		num := rand.Intn(100)
		fmt.Println("生产者生产了: ", num)
		// 往管道中写入数据
		myCh<-num
		//time.Sleep(time.Millisecond * 500)
	}
	// 生产完毕之后关闭管道
	close(myCh)
	fmt.Println("生产者停止生产")
}
// 定义消费者
func consumer4(myCh <-chan int)  {
	// 不断从管道中获取数据, 直到管道关闭位置
	for{
		if num, ok := <-myCh; !ok{
			break
		}else{
			fmt.Println("---消费者消费了", num)
		}
	}
	fmt.Println("消费者停止消费")

}
