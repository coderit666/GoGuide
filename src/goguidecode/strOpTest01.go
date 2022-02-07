package main

import (
	"fmt"
	"strings"
)

func main() {
	lenTest()
	fmt.Println()
	findTest()
	containsTest()
	compareTest()
	toSomethingTest()
	splitAndJoinTest()
	trimTest()
}

func lenTest() {
	str1 := "lnj"
	fmt.Println(len(str1)) // 3
	str2 := "公号:代码情缘"
	fmt.Println(len(str2)) // 19

	str := "公号：代码情缘"
	// 注意byte占1个字节, 只能保存字符不能保存汉字,因为一个汉字占用3个字节
	arr1 := []byte(str) // 21
	fmt.Println(len(arr1))
	for _, v := range arr1{
		fmt.Printf("%c", v) // å¬å·ï¼ä»£ç æç¼7
	}

	// Go语言中rune类型就是专门用于保存汉字的
	arr2 := []rune(str)
	fmt.Println(len(arr2)) // 7
	for _, v := range arr2{
		fmt.Printf("%c", v) // 公号：代码情缘
	}
}

func findTest() {
	// 查找`字符`在字符串中第一次出现的位置, 找不到返回-1
	res := strings.IndexByte("hello 李南江", 'l')
	fmt.Println(res) // 2

	// '李' is not byte, panic. # command-line-arguments constant 26446 overflows byte
	//res = strings.IndexByte("hello 李南江", '李')

	// 查找`汉字`OR`字符`在字符串中第一次出现的位置, 找不到返回-1
	res = strings.IndexRune("hello 李南江", '李')
	fmt.Println(res) // 6
	res = strings.IndexRune("hello 李南江", 'l')
	fmt.Println(res) // 2

	// 查找`汉字`OR`字符`中任意一个在字符串中第一次出现的位置, 找不到返回-1
	res = strings.IndexAny("hello 李南江", "wml")
	fmt.Println(res) // 2
	// 会把wmhl拆开逐个查找, w、m、h、l只要任意一个被找到, 立刻停止查找
	res = strings.IndexAny("hello 李南江", "wmhl")
	fmt.Println(res) // 0
	// 查找`子串`在字符串第一次出现的位置, 找不到返回-1
	res = strings.Index("hello 李南江", "llo")
	fmt.Println(res) // 2
	// 会把lle当做一个整体去查找, 而不是拆开
	res = strings.Index("hello 李南江", "lle")
	fmt.Println(res) // -1
	// 可以查找字符也可以查找汉字
	res = strings.Index("hello 李南江", "李")
	fmt.Println(res) // 6

	// 会将字符串先转换为[]rune, 然后遍历rune切片逐个取出传给自定义函数
	// 只要函数返回true,代表符合我们的需求, 既立即停止查找
	res = strings.IndexFunc("hello 李南江", custom)
	fmt.Println(res) // 4

	// 倒序查找`子串`在字符串第一次出现的位置, 找不到返回-1
	res = strings.LastIndex("hello 李南江", "l")
	fmt.Println(res) // 3
}

func custom(r rune) bool {
	fmt.Printf("被调用了, 当前传入的是%c\n", r)
	if r == 'o' {
		return true
	}
	return false
}

//判断字符串中是否包含子串
func containsTest() {
	// 查找`子串`在字符串中是否存在, 存在返回true, 不存在返回false
	// 底层实现就是调用strings.Index函数
	res := strings.Contains( "hello 代码情缘", "llo")
	fmt.Println(res) // true

	// 查找`汉字`OR`字符`在字符串中是否存在, 存在返回true, 不存在返回false
	// 底层实现就是调用strings.IndexRune函数
	res = strings.ContainsRune( "hello 代码情缘", 'l')
	fmt.Println(res) // true
	res = strings.ContainsRune( "hello 代码情缘", '李')
	fmt.Println(res) // false

	// 查找`汉字`OR`字符`中任意一个在字符串中是否存在, 存在返回true, 不存在返回false
	// 底层实现就是调用strings.IndexAny函数
	res = strings.ContainsAny( "hello 代码情缘", "wmhl")
	fmt.Println(res) // true

	// 判断字符串是否以某个字符串开头
	res = strings.HasPrefix("lnj-book.avi", "lnj")
	fmt.Println(res) // true

	// 判断字符串是否已某个字符串结尾
	res = strings.HasSuffix("lnj-book.avi", ".avi")
	fmt.Println(res) // true
}

func compareTest() {
	// 比较两个字符串大小, 会逐个字符地进行比较ASCII值
	// 第一个参数 >  第二个参数 返回 1
	// 第一个参数 <  第二个参数 返回 -1
	// 第一个参数 == 第二个参数 返回 0
	res := strings.Compare("bcd", "abc")
	fmt.Println(res) // 1
	res = strings.Compare("bcd", "bdc")
	fmt.Println(res) // -1
	res = strings.Compare("bcd", "bcd")
	fmt.Println(res) // 0

	// 判断两个字符串是否相等, 可以判断字符和中文
	// 判断时会忽略大小写进行判断
	res2 := strings.EqualFold("abc", "def")
	fmt.Println(res2) // false
	res2 = strings.EqualFold("abc", "abc")
	fmt.Println(res2) // true
	res2 = strings.EqualFold("abc", "ABC")
	fmt.Println(res2) // true
	res2 = strings.EqualFold("代码情缘", "代码情缘")
	fmt.Println(res2) // true
}

//字符串转换
func toSomethingTest() {
	// 将字符串转换为小写
	res := strings.ToLower("ABC")
	fmt.Println(res) // abc

	// 将字符串转换为大写
	res = strings.ToUpper("abc")
	fmt.Println(res) // ABC

	// 将字符串转换为标题格式, 大部分`字符`标题格式就是大写
	res = strings.ToTitle("hello world,./-='!")
	fmt.Println(res) // HELLO WORLD
	res = strings.ToTitle("HELLO WORLD")
	fmt.Println(res) // HELLO WORLD

	// 将单词首字母变为大写, 其它字符不变
	// 单词之间用空格OR特殊字符隔开
	res = strings.Title("hello world")
	fmt.Println(res) // Hello World
}

func splitAndJoinTest() {
	// 按照指定字符串切割原字符串
	// 用,切割字符串
	arr1 := strings.Split("a,b,c", ",")
	fmt.Println(arr1) // [a b c]
	arr2 := strings.Split("ambmc", "m")
	fmt.Println(arr2) // [a b c]

	// 按照指定字符串切割原字符串, 并且指定切割为几份,从前开始尽量切分到n-1份
	// 如果最后一个参数为0, 那么会范围一个空数组
	arr3 := strings.SplitN("a,b,c,d", ",", 3)
	fmt.Println(arr3) // [a b c,d]
	arr4 := strings.SplitN("a,b,c", ",", 0)
	fmt.Println(arr4) // []

	// 按照指定字符串切割原字符串, 切割时包含指定字符串
	arr5 := strings.SplitAfter("a,b,c", ",")
	fmt.Println(arr5) // [a, b, c]

	// 按照指定字符串切割原字符串, 切割时包含指定字符串, 并且指定切割为几份
	arr6 := strings.SplitAfterN("a,b,c", ",", 2)
	fmt.Println(arr6) // [a, b,c]

	// 按照空格切割字符串, 多个空格会合并为一个空格处理
	arr7 := strings.Fields("a  b c    d")
	fmt.Println(arr7) // [a b c d]

	// 将字符串转换成切片传递给函数之后由函数决定如何切割
	// 类似于IndexFunc
	arr8 := strings.FieldsFunc("a,b,c", fieldsTest)
	fmt.Println(arr8) // [a b c]

	// 将字符串切片按照指定连接符号转换为字符串
	sce := []string{"aa", "bb", "cc"}
	str1 := strings.Join(sce, "-")
	fmt.Println(str1) // aa-bb-cc


	// 返回count个s串联的指定字符串
	str2 := strings.Repeat("abc", 2)
	fmt.Println(str2) // abcabc

	// 第一个参数: 需要替换的字符串
	// 第二个参数: 旧字符串
	// 第三个参数: 新字符串
	// 第四个参数: 用新字符串 替换 多少个旧字符串
	// 注意点: 传入-1代表只要有旧字符串就替换. 传入大于0的数表示从前到后替换的个数。
	// 注意点: 替换之后会生成新字符串, 原字符串不会受到影响
	str3 := "abcdefabcdefabc"
	str4 := strings.Replace(str3, "abc", "mmm", -1)
	fmt.Println(str3) // abcdefabcdefabc
	fmt.Println(str4) // mmmdefmmmdefmmm
}

func fieldsTest(r rune) bool {
	fmt.Printf("被调用了, 当前传入的是%c\n", r)
	if r == ',' {
		return true
	}
	return false
}

func trimTest() {
	// 去除字符串两端指定字符
	str1 := strings.Trim("!!!abc!!!def!!!", "!")
	fmt.Println(str1) // abc!!!def
	// 去除字符串左端指定字符
	str2 := strings.TrimLeft("!!!abc!!!def!!!", "!")
	fmt.Println(str2) // abc!!!def!!!
	// 去除字符串右端指定字符
	str3 := strings.TrimRight("!!!abc!!!def!!!", "!")
	fmt.Println(str3) // !!!abc!!!def
	// 去除字符串两端空格
	str4 := strings.TrimSpace("   abc!!!def ")
	fmt.Println(str4) // abc!!!def

	// 按照方法定义规则,去除字符串两端符合规则内容
	str5 := strings.TrimFunc("!!!abc!!!def!!!", trimFunc)
	fmt.Println(str5) // abc!!!def
	// 按照方法定义规则,去除字符串左端符合规则内容
	str6 := strings.TrimLeftFunc("!!!abc!!!def!!!", trimFunc)
	fmt.Println(str6) // abc!!!def!!!
	//  按照方法定义规则,去除字符串右端符合规则内容
	str7 := strings.TrimRightFunc("!!!abc!!!def!!!", trimFunc)
	fmt.Println(str7) // !!!abc!!!def

	// 取出字符串开头的指定字符串
	str8 := strings.TrimPrefix("lnj-book.avi", "lnj-")
	fmt.Println(str8) // book.avi

	// 取出字符串结尾的指定字符串
	str9 := strings.TrimSuffix("lnj-book.avi", ".avi")
	fmt.Println(str9) // lnj-book
}

func trimFunc(r rune) bool {
	fmt.Printf("被调用了, 当前传入的是%c\n", r)
	if r == '!' {
		return true
	}
	return false
}