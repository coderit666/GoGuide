package main

import (
	"fmt"
	"runtime"
	"time"
)

func dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("dancing")
		time.Sleep(time.Millisecond)
		runtime.Gosched()
	}
}

func sing(){
	for i := 0; i < 5; i++ {
		fmt.Println("singing")
		time.Sleep(time.Millisecond)
		runtime.Gosched()
	}
}

func test()  {
	fmt.Println("test do go exit")
	// 只会结束当前函数, 协程中的其它代码会继续执行
	//return
	// 会结束整个协程, Goexit之后整个协程中的其它代码不会执行
	runtime.Goexit()
	fmt.Println("not reach")
}

func main() {
	go sing()
	go dance()
	go func() {
		fmt.Println("abc")
		test()
		fmt.Println("test exit this not print")
	}()

	for{
		;
	}
}

