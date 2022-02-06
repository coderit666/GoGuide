package main
import "fmt"

func div(a, b int) (res int) {
	// 定义一个延迟调用的函数, 用于捕获panic异常
	// 注意: 一定要在panic之前定义
	defer func() {
		if err := recover(); err != nil{
			res = -1
			fmt.Println(err) // 除数不能为0
		}
	}()

	if b == 0 {
		//err = errors.New("除数不能为0")
		panic("除数不能为0")
	}else{
		res = a / b
	}

	defer func() {
		panic("异常被defer捕获")
	}()

	return
}

func setValue(arr []int, index int ,value int)  {
	arr[index] = value
}

func main() {
	res := div(10, 0)
	fmt.Println(res) // -1

	defer func() {
		panic("异常被defer捕获")
	}()

	//panic("异常2")
}

func test2()  {
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