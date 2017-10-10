package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

/*
考察： defer执行
结果：
10, 1, 2, 3 -> 3
20, 0, 2, 2 -> 2
2, 0, 2, 2 -> 2
1, 1, 3, 4 -> 4
解答：
	1. defer语句是做的压栈操作， 将当前涉及的变量压入堆栈；
	2. 在函数return时， 执行出栈操作；
*/

func main(){
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
