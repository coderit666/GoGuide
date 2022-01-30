package main

import "fmt"

func main() {
    // 1.定义一个普通变量
    var num = 666
    // 2.定义一个指针变量
    var p *int = &num
    fmt.Printf("%p\n", &num) // 0xc0000b4008
    fmt.Printf("%p\n", p)    // 0xc0000b4008
    fmt.Printf("%T\n", p)    // *int
    // 3.通过指针变量操作指向的存储空间
    *p = 888
    // 4.指针变量操作的就是指向变量的存储空间
    fmt.Println(num) // 888
    fmt.Println(*p)  // 888
    
    var arr = [3]int{1, 2, 3}
    // 1.错误, 在Go语言中必须类型一模一样才能赋值
    // arr类型是[3]int, p1的类型是*[3]int
    var p1 *[3]int
    fmt.Printf("%T\n", arr)
    fmt.Printf("%T\n", p1)
    fmt.Printf("%T\n", &arr)
    // p1 = arr // 报错 type(arr) != type(p1)
    p1 = &arr // 正确 type(&arr) = type(p1)
    p1[1] = 6 // 我猜这个是个语法糖，和p6结构体的指针一样？
    (*p1)[2] = 8
    fmt.Println(arr[1])
    
    // 2 &arr[0]的类型是*int, p3作为普通指针指向 arr[0]
    fmt.Printf("%T\n", &arr[0])
    p3 := &arr[0]
    *p3 = 6
    
    // 1.定义一个切片
    var sce = []int{1, 3, 5}
    // 2.打印切片的地址
    // 切片变量中保存的地址, 也就是指向的那个数组的地址 sce = 0xc0420620a0
    fmt.Printf("sce = %p\n", sce)
    fmt.Println(sce) // [1 3 5]
    // 切片变量自己的地址, &sce = 0xc04205e3e0
    fmt.Printf("&sce = %p\n", &sce)
    fmt.Println(&sce) // &[1 3 5]
    // 3.定义一个指向切片的指针
    // 因为必须类型一致才能赋值, 所以将切片变量自己的地址给了指针
    var p4 *[]int
    p4 = &sce
    // 4.打印指针保存的地址
    // 直接打印p打印出来的是保存的切片变量的地址 p = 0xc04205e3e0
    fmt.Printf("p4 = %p\n", p4)
    fmt.Println(p4) // &[1 3 5]
    // 打印*p打印出来的是切片变量保存的地址, 也就是数组的地址 *p = 0xc0420620a0
    fmt.Printf("*p4 = %p\n", *p4)
    fmt.Printf("p4 type %T\n", p4)
    fmt.Printf("*p4 type %T\n", *p4)
    fmt.Println(*p4) // [1 3 5]
    // 5.修改切片的值
    // 通过*p找到切片变量指向的存储空间(数组), 然后修改数组中保存的数据
    (*p4)[1] = 666
    fmt.Println(sce)
    
    // 字典指针与切片指针很像
    var dict map[string]string = map[string]string{"name": "lnj", "age": "33"}
    var p5 *map[string]string = &dict
    fmt.Printf("%T\n", dict)
    fmt.Printf("%T\n", p5)
    fmt.Printf("%T\n", &p5)
    var pp5 = &p5
    (*(*pp5))["name"] = "?"
    (*p5)["name"] = "zs"
    fmt.Println(dict)
    
    type Student struct {
        name string
        age  int
    }
    var p6 = &Student{}
    // 方式一: 传统方式操作
    // 修改结构体中某个属性对应的值
    // 注意: 由于.运算符优先级比*高, 所以一定要加上()
    (*p6).name = "lnj"
    // 获取结构体中某个属性对应的值
    fmt.Println((*p6).name) // lnj
    
    // 方式二: 通过Go语法糖操作
    // Go语言作者为了程序员使用起来更加方便, 在操作指向结构体的指针时可以像操作接头体变量一样通过.来操作
    // 编译时底层会自动转发为(*p).age方式
    p6.age = 33
    fmt.Println(p6.age) // 33
    
    //指针作为函数参数和返回值
    //如果指针类型作为函数参数, 修改形参会影响实参
    //不能将函数内的指向局部变量的指针作为返回值, 函数结束指向空间会被释放
    //可以将函数内的局部变量作为返回值, 本质是拷贝一份
}
