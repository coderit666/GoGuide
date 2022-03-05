package main

import "fmt"

func main() {
	// map格式:var dic map[key数据类型]value数据类型
	var dic = map[int]int{0: 1, 1: 3, 3: 5}
	fmt.Println(dic, len(dic)) // map[0:1 1:3 3:5] 3
	dic[4] = 7
	fmt.Println(dic, len(dic)) // map[0:1 1:3 3:5 4:7] 4
	delete(dic, dic[0])
	fmt.Println(dic, len(dic)) // map[0:1 3:5 4:7] 3

	// 查询: 通过ok-idiom模式判断指定键值对是否存储 idiom(习语、俗语)
	// i: 0 1 2 3 4
	// v: 1 0 0 5 7
	// ok: true false false true true
	for i := 0; i < 5; {
		v, ok := dic[i]
		fmt.Println(i, v, ok)
		i++
	}
	// k: 0 3 4
	// v: 1 5 7
	for k, v := range dic {
		fmt.Println(k, v)
	}
}
