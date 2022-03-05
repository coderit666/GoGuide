package main

import (
	"fmt"
)

func main() {

	// generate slice from array. array[startIndex:endIndex]
	var arr = [5]int{0, 1, 2, 3, 4}
	// var sce = arr[a:b] 从索引下标a开始，到第b个元素结束。 最终取到的索引范围 [a, b - 1]
	var sce = arr[1:4]
	fmt.Println(sce) // [1 2 3]
	// len = 切片结束位置索引 - 开始位置 3
	// cap = 切片截取的数组长度 - 切片开始位置 4
	fmt.Println("sce len =", len(sce), "cap = ", cap(sce))
	// 数组地址就是数组首元素的地址
	fmt.Printf("%p\n", &arr)    // 0xc00001c0f0
	fmt.Printf("%p\n", &arr[1]) // 0xc00001c0f8
	// 切片地址就是数组中指定的开始元素的地址. arr[1:4]开始地址为1, 所以就是arr[1]的地址
	fmt.Printf("%p\n", sce) // 0xc00001c0f8
	// 同理 &arr[2] == &sce[1]
	fmt.Printf("%p\n", &arr[2])                         // 0xc00001c100
	fmt.Printf("%p\n", &sce[1])                         // 0xc00001c100
	fmt.Println(&arr[2] == &sce[1] && arr[2] == sce[1]) // true

	// 目前slice长度为3 索引最大为2 sce[3] = 68 会报错 -- panic: runtime error: index out of range [3] with length 3
	// sce[3] = 68

	// 未超过cap 不会扩容
	sce = append(sce, 67)
	// 扩容1后 指定索引位置3 进行赋值. slice是指针操作，原数组arr的内容跟着发生改动
	sce[3] = 68
	// len = 4 cap =  4
	fmt.Println("sce len =", len(sce), "cap = ", cap(sce))
	// arr =  [0 1 2 3 68] sce = [1 2 3 68]
	fmt.Println("arr = ", arr, "sce =", sce)

	sce = append(sce, 99)
	// 扩容1后，len > cap, slice 扩容一倍。len = 5 cap =  8
	fmt.Println("sce len =", len(sce), "cap = ", cap(sce))

	// 通过make函数创建 make(类型, 长度, 容量)
	// 内部会先创建一个数组, 然后让切片指向数组
	// 如果没有指定容量,那么容量和长度一样
	var sce1 = []int{1, 3, 5}
	var sce2 = make([]int, 5)
	fmt.Println("len =", len(sce1), "cap = ", cap(sce1)) // len = 3 cap =  3
	fmt.Println("len =", len(sce2), "cap = ", cap(sce2)) // len = 5 cap =  5
	fmt.Printf("赋值前:%p\n", sce1)                         // 0xc0000c2020
	fmt.Printf("赋值前:%p\n", sce2)                         // 0xc0000b6090
	// 将sce2的指向修改为sce1, 此时sce1和sce2底层指向同一个数组
	sce2 = sce1
	fmt.Printf("赋值后:%p\n", sce1)                         // 0xc0000c2020
	fmt.Printf("赋值后:%p\n", sce2)                         // 0xc0000c2020
	fmt.Println("len =", len(sce1), "cap = ", cap(sce1)) // len = 3 cap =  3
	fmt.Println("len =", len(sce2), "cap = ", cap(sce2)) // len = 3 cap =  3
	fmt.Println(sce1, sce2)                              // [1 3 5] [1 3 5]

	//copy(sce2, sce1)
	sce2 = make([]int, 5)
	copy(sce2, sce1)
	// sce2 已经是独立的地址\堆内存
	fmt.Printf("赋值后:%p\n", sce1)                         // 0xc0000c2020
	fmt.Printf("赋值后:%p\n", sce2)                         // 0xc00001c180
	fmt.Println("len =", len(sce1), "cap = ", cap(sce1)) // len = 3 cap =  3
	fmt.Println("len =", len(sce2), "cap = ", cap(sce2)) // len = 5 cap =  5
	sce2[1] = 666
	fmt.Println(sce1, sce2) // [1 3 5] [1 666 5 0 0 ]

	var sce3 []int
	fmt.Println("sce3 len =", len(sce3), "cap = ", cap(sce3)) // sce3 len = 0 cap =  0
	sce3 = make([]int, 4)
	sce3[0] = 1
	fmt.Println("sce3 len =", len(sce3), "cap = ", cap(sce3)) // sce3 len = 4 cap =  4
	sce3 = []int{3, 2, 0, 0, 0, 0}
	fmt.Println("sce3 len =", len(sce3), "cap = ", cap(sce3)) // sce3 len = 6 cap =  6

	// 字符串的底层是[]byte数组, 所以字符也支持切片相关操作
	var str string
	str = "hello-world"
	// 通过字符串切片 生成新字符串
	sce4 := str[3:]
	fmt.Println(sce4) // lo-world

	sce5 := make([]byte, 10)
	// 将字符串拷贝到切片中
	// 第二个参数只能是slice或者是数组
	copy(sce5, str)
	fmt.Println(sce5) //[97 98 99 100 101 102 103 0 0 0]
}