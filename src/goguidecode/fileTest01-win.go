package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	/*
		Create函数
		func Create(name string) (file *File, err error)
		Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件
		如果文件存在会覆盖原有文件
		Write函数
		func (f *File) Write(b []byte) (n int, err error)
		将指定字节数组写入到文件中
		WriteString函数
		func (f *File) WriteString(s string) (ret int, err error)
		将指定字符串写入到文件中
	*/

	filePath := "d:/tmpTestCreate.txt"
	// 1.创建一个文件
	fp, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}
	// 2.关闭打开的文件
	defer func() {
		err := fp.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// 3.往文件中写入数据
	// 注意: Windows换行是\r\n
	bytes := []byte{'c', 'i', 't', 'y', '\r', '\n'}
	fp.Write(bytes)

	fp.WriteString("www.baidu.com\r\n")
	fp.WriteString("www.kuaishou.com\r\n")
	// 注意: Go语言采用UTF-8编码, 一个中文占用3个字节
	fp.WriteString("城市\n")


	// 注意点: 第三个参数在Windows没有效果
	// -rw-rw-rw- (666)   所有用户都有文件读、写权限。
	//-rwxrwxrwx (777)  所有用户都有读、写、执行权限。
	// 4.打开文件
	//fp, err := os.OpenFile("d:/lnj.txt", os.O_CREATE|os.O_RDWR, 0666)
	fp, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	// 注意点:
	// 如果O_RDWR模式打开, 被打开文件已经有内容, 会从最前面开始覆盖
	// 如果O_APPEND模式打开, 被打开文件已经有内容, 会从在最后追加
	// 5.往文件中写入数据
	bytes = []byte{'2','c','i','t','y','\r','\n'}
	fp.Write(bytes)
	fp.WriteString("www.google.com\r\n")

	// 6.1创建缓冲区
	w := bufio.NewWriter(fp)

	// 6.2.写入数据到缓冲区
	bytes = []byte{'3','c','i','t','y','\r','\n'}
	w.Write(bytes)
	w.WriteString("www.tencent.com\r\n")

	// 6.3.将缓冲区中的数据刷新到文件
	w.Flush()

	// 7.写入数据到指定文件
	str2 := "2022BeijingWinterOlympicGames\r\n"
	data := make([]byte, len(str2))
	copy(data, str2)
	//data := []byte(str2)  直接將str强转为byte切片
	err = ioutil.WriteFile("d:/tempTest02.txt", data, 0666)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("写入成功")
	}

	info, err := os.Stat("d:/notExist.txt")
	if err == nil {
		fmt.Println("文件存在")
		fmt.Println(info.Name())
	}else if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	}else{
		fmt.Println("不确定")
	}

	info, err = os.Stat(filePath)
	if err == nil {
		fmt.Println("文件存在")
		fmt.Println(info.Name())
		fmt.Println(info.Mode())
	}else if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	}else{
		fmt.Println("不确定")
	}
}
