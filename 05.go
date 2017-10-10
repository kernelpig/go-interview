package main

import (
	"runtime"
	"fmt"
)

/*
考察： select随机性
结果： 可能触发异常， 也可能不会
其他： select在多个chan间等待执行， 三个原则：
	1. select中只要有一个case能return， 则立即执行；
	2. 当如果同一时间有多个case均能return则伪随机任意选取一个；
	3. 如果没有一个case能return则执行default；
*/
func main() {
	runtime.GOMAXPROCS(1)

	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	int_chan <- 1
	string_chan <- "hello"

	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
