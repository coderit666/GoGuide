package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var t time.Time = time.Now()
	fmt.Println(t) // 2022-02-09 00:57:12.08913 +0800 CST m=+0.010149201

	fmt.Println(t.Year()) // 2022
	fmt.Println(t.Month()) // February
	fmt.Println(t.Day()) // 9
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())

	fmt.Printf("当前的时间是: %d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	var dateStr = fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println("当前的时间是:", dateStr)

	// 2006/01/02 15:04:05这个字符串是Go语言规定的, 各个数字都是固定的. 06年1月2日3点4分5秒（下午）
	// 字符串中的各个数字可以只有组合, 这样就能按照需求返回格式化好的时间
	str1 := t.Format("2006/01/02 15:04:05")
	fmt.Println(str1)
	strx := t.Format("2017-09-07 18:05:32")
	fmt.Println(strx)
	str2 := t.Format("2006-01-02")
	fmt.Println(str2)
	str3 := t.Format("15:04:05")
	fmt.Println(str3)

	for i := 0; i < 2;{
		// 1秒钟打印一次
		time.Sleep(time.Second * 1)
		fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "Hello one second")
		i++
	}

	t = time.Now()
	fmt.Println(t.Nanosecond())
	for i := 0; i < 2; {
		// 0.1秒打印一次
		time.Sleep(time.Millisecond * 100)
		t := time.Now()
		fmt.Println(t.Nanosecond(), "Hello 100 Millisecond")
		i++
	}

	// 获取1970年1月1日距离当前的时间(秒)
	fmt.Println(t.Unix())
	// 获取1970年1月1日距离当前的时间(毫秒)
	fmt.Println(t.UnixMilli())
	// 获取1970年1月1日距离当前的时间(纳秒)
	fmt.Println(t.UnixNano())

	// 创建随机数种子
	rand.Seed(time.Now().UnixNano())
	// 生成一个随机数
	fmt.Println(rand.Intn(10))

}
