package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var lock sync.Mutex

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

	go person1()
	go person2()

	for{
		;
	}
}

func goToilet(str string)  {
	// 让先来的人拿到锁, 把当前函数锁住, 其它人都无法执行
	// 上厕所关门
	lock.Lock()
	for _, v := range str{
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("%c", v)
	}
	fmt.Println(" is shitting")
	// 先来的人执行完毕之后, 把锁释放掉, 让其它人可以继续使用当前函数
	// 上厕所开门
	lock.Unlock()
}

func person1()  {
	goToilet("zhangSan")
}

func person2()  {
	goToilet("LiSi")
}