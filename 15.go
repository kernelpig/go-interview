package main

import "fmt"

/*
结果： ./15.go:7:15: first argument to append must be slice; have *[]int
解答： new用于分配内存返回指针，应使用make
*/
func main(){
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}
