package main

import "fmt"

func GetValue() int {
	return 1
}

/*
结果： ./12.go:12:2: cannot type switch on non-interface value i (type int)
解答：类型断言只能用于interface类型， 具体的类型不可以
*/
func main() {
	i := GetValue()

	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
