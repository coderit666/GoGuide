package main
import "fmt"

func test3(a int, b int)  {
	// 如果有异常写在defer中, 并且其它异常写在defer后面, 那么只有defer中的异常会被捕获
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err) // 异常A
		}
	}()

	defer func() {
		panic("异常B")
	}()
	panic("异常A")
}

func test1()  {
	// 多个异常,只有第一个会被捕获
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err) // 异常A
		}
	}()
	panic("异常A") // 相当于return, 后面代码不会继续执行
	panic("异常B")
}

func main() {
	test3(10, 0)
}