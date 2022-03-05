package main

import "fmt"

type aer interface {
	fna()
}

type ber interface {
	aer
	fnb()
}

type studier interface {
	read()
}

// Person Person实现了超集接口所有方法
type Person struct{
	name string
	age int
}

func (p Person) fna() {
	fmt.Println("Person 实现A接口中的方法")
}
func (p Person) fnb() {
	fmt.Println("Person 实现B接口中的方法")
}

func (p Person)read()  {
	fmt.Println(p.name, "正在学习")
}

// Student Student实现了子集接口所有方法
type Student struct{}

func (s Student) fna() {
	fmt.Println("Student 实现A接口中的方法")
}

// 1.定义一个接口
type usber interface {
	start()
	stop()
}

// 2.自定义int类型
type integer int

// 2.实现接口中的所有方法
func (i integer) start() {
	fmt.Println("int类型实现 start接口", i)
}
func (i integer) stop() {
	fmt.Println("int类型实现 stop接口", i)
}

func main() {
	var i ber
	// 子集接口变量不能转换为超集接口变量
	//i = Student{}
	fmt.Println(i)
	var j1 aer
	// 超集接口变量可以自动转换成子集接口变量,
	j1 = Person{}
	j1.fna()

	var j2 ber
	j2 = Person{}
	j2.fna()
	j2.fnb()

	var j3 aer
	j3 = Student{}
	j3.fna()

	var ii integer
	ii.start() // int类型实现接口
	ii.stop()  // int类型实现接口

	var s studier
	s = Person{"lnj", 33}

	// 2.利用ok-idiom模式将接口类型还原为原始类型
	// s.(Person)这种格式我们称之为: 类型断言
	if p, ok := s.(Person); ok {
		p.name = "zs"
		p.age = 32
		fmt.Println(p)
	}

	// 2.通过 type switch将接口类型还原为原始类型
	// 注意: type switch不支持fallthrought
	switch p := s.(type) {
	case Person:
		p.name = "ppq"
		p.age = 56
		fmt.Println(p) // {zs 33}
	default:
		fmt.Println("不是Person类型")
	}
}
