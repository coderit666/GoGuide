package main

import "fmt"

func main() {

	// producerAndConsumerNaiveTest01.go
	/*
		goRoutineInGoTest01.go 实现并发的代码中为了保持主线程不挂掉, 我们都会在最后写上一个死循环或者写上一个定时器来实现等待goroutine执行完毕
		producerAndConsumerNaiveTest01.go 实现并发的代码中为了解决生产者消费者资源同步问题, 我们利用加锁来解决, 但是这仅仅是一对一的情况, 如果是一对多或者多对多, 上述代码还是会出现问题
		综上所述, 企业开发中需要一种更牛X的技术来解决上述问题, 那就是管道(Channel)
	*/

	// 1.声明一个管道
	var mych chan int
	// 2.初始化一个管道
	mych = make(chan int, 4)
	// 3.查看管道的长度和容量
	fmt.Println("长度是", len(mych), "容量是", cap(mych))
	// 4.像管道中写入数据
	mych <- 777
	mych <- 666
	fmt.Println("长度是", len(mych), "容量是", cap(mych))
	// 5.取出管道中写入的数据
	num := <-mych
	fmt.Println("num = ", num)
	fmt.Println("长度是", len(mych), "容量是", cap(mych))

	// 注意点: 管道中只能存放声明的数据类型, 不能存放其它数据类型
	//mych<-3.14

	// 注意点: 管道中如果已经没有数据,
	// 并且检测不到有其它协程再往管道中写入数据, 那么再取就会报错
	//num = <-mych
	//fmt.Println("num = ", num)

	mych <- 777
	mych <- 888
	mych <- 999

	// 3.遍历管道
	// 第一次遍历i等于0, len = 3
	// 第二次遍历i等于1, len = 2
	// 第三次遍历i等于2, len = 1
	for i := 0; i < len(mych); i++ {
		fmt.Print("--len--", len(mych), " i ", i, " --")
		fmt.Println(<-mych) // 输出结果不正确
	}

	// 3.写入完数据之后先关闭管道
	// 注意点: 管道关闭之后只能读不能写
	close(mych)
	//mych <- 999 // 报错

	// 4.遍历管道
	// 利用for range遍历, 必须先关闭管道, 否则会报错
	//for value := range mych{
	//	fmt.Println(value)
	//}

	// close主要用途:
	// 在企业开发中我们可能不确定管道有还没有有数据, 所以我们可能一直获取
	// 但是我们可以通过ok-idiom模式判断管道是否关闭, 如果关闭会返回false给ok. 必须先关闭管道, 否则会报错
	for {
		if num, ok := <-mych; ok {
			fmt.Println(num)
		} else {
			break
		}
	}
	fmt.Println("数据读取完毕")

	/*
		注意点: 如果管道中数据已满, 再写入就会报错
		mych<- 888
		mych<- 999
		goroutine 1 [chan send]:
		main.main()
		D:/Devops/gohome/GoGuide/src/goguidecode/chanTest01.go:40 +0x3cf
	*/

	var myCh2 = make(chan int, 0)
	// 无缓冲管道
	// 只有两端同时准备好才不会报错
	go func() {
		fmt.Println(<-myCh2)
	}()
	// 只写入, 不读取会报错
	myCh2 <- 1
	fmt.Println("len =", len(myCh2), "cap =", cap(myCh2))
	// 写入之后在同一个线程读取也会报错
	//fmt.Println(<-myCh2)

	// 在主程中先写入, 在子程中后读取也会报错
	//go func() {
	//	fmt.Println(<-myCh2)
	//}()

	// 1.定义一个双向管道
	var myCh chan int = make(chan int, 5)
	// 2.将双向管道转换单向管道
	var myCh4 chan<- int
	myCh4 = myCh
	fmt.Println(myCh4)
	var myCh3 <-chan int
	myCh3 = myCh
	fmt.Println(myCh3)
	// 3.双向管道,可读可写
	myCh <- 1
	fmt.Println(<-myCh)
	// 3.只写管道,只能写, 不能读
	myCh4 <- 666
	//	fmt.Println(<-myCh2)
	// 4.指读管道, 只能读,不能写
	fmt.Println(<-myCh3)
	//myCh3<-666
	// 注意点: 管道之间赋值是地址传递, 以上三个管道底层指向相同容器
}
