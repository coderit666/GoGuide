package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func main() {
    // 1.打开一个文件
    // 注意: 文件不存在不会创建, 会报错
    // 注意: 通过Open打开只能读取, 不能写入
    /* 文件内容
        helloFile
        helloFile1
        helloFile2
     */
    fp, err := os.Open("/Users/citi/testFile01.txt")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(fp)
    }
    
    defer func() {
        err = fp.Close()
        if err != nil {
            fmt.Println(err)
        }
    }()
    
    // 3.读取指定字节个数据
    // 注意点: \n也会被读取进来
    buf := make([]byte, 5)
    count, err := fp.Read(buf)
    if err != nil {
    	fmt.Println(err) // hello
    }else{
    	fmt.Println(count) // 5
    	fmt.Println(string(buf))
    }
    
    // 4.循环使用一个切片，读取文件中所有内容, 直到文件末尾为止
    buf = make([]byte, 7)
    for{
        count, err := fp.Read(buf)
        // 注意: 这行代码要放到判断EOF之前, 否则会出现少读一行情况
        fmt.Print(string(buf[:count]))
        /* 3 中读走了5个byte了
            File
            helloFile1
            helloFile2
        */
        if err == io.EOF {
            break
        }
    }
    
    // 重新打开一遍。不重开，就在文件末尾。
    fp, err = os.Open("/Users/citi/testFile01.txt")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(fp)
    }
    // 3.读取一行数据
    // 创建读取缓冲区, 默认大小4096
    r1 :=bufio.NewReader(fp)
    buf21, err := r1.ReadBytes('\n')
    buf22, err := r1.ReadString('\n')
    if err != nil{
    	fmt.Println(err)
    }else{
    	fmt.Print(string(buf21))
        fmt.Print(buf22)
    }
    
    // 新建一个缓冲区，有问题，读取不到 r1 :=bufio.NewReader(fp)
    // 4.读取文件中所有内容, 直到文件末尾为止
    for{
        //buf, err := r.ReadBytes('\n')
        buf, err := r1.ReadString('\n')
        fmt.Println(buf)
        if err == io.EOF{
            break
        }
    }
    
    buf, err = ioutil.ReadFile("/Users/citi/testFile01.txt")
    if err !=nil {
        fmt.Println(err)
    }else{
        fmt.Println(string(buf))
    }
    
    
}