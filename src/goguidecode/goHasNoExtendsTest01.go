package main

import "fmt"

type person struct {
	name string
	age int
}

func (p person) say()  {
	fmt.Println("name is", p.name, "age is", p.age)
}

type student struct {
	person
	score float32
}

func (s student) say() {
	fmt.Println("name is", s.name, "age is", s.age, "score is ", s.score)
}

// Animal 1.定义接口
type Animal interface {
	Eat()
}
type Dog struct {
	name string
	age int
}

// Eat 2.实现接口方法
func (d Dog)Eat()  {
	fmt.Println(d.name, "正在吃东西")
}

type Cat struct {
	name string
	age int
}

// Eat 2.实现接口方法
func (c Cat)Eat()  {
	fmt.Println(c.name, "正在吃东西")
}

// Special 3.对象特有方法
func (c Cat)Special()  {
	fmt.Println(c.name, "特有方法")
}

func main() {
	stu := student{person{"zs", 18}, 59.9}
	stu.say() // name is zs age is 18 score is  59.9
	stu.person.say() // name is zs age is 18

	var a Animal

	a = Dog{"dog", 9}
	a.Eat()

	a = Cat{"cat", 10}
	a.Eat()

	if cat,ok := a.(Cat);ok {
		fmt.Println("ok-idiom 命中")
		cat.Special()
	}

	switch scat := a.(type) {
	case Cat:
		fmt.Println("switch type 命中")
		scat.Special()
	default:
		fmt.Println("not Cat type")
	}

}
