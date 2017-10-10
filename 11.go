package main

import "fmt"

type People11 interface {
	Show()
}

type Student11 struct {

}

func (stu *Student11) Show() {

}

func live() People11 {
	/*
	所有最好明确返回nil，不要类型转换
	*/
	var stu *Student11
	return stu
}

/*
结果： BBB
解答： 空接口与空指针的区别， 空接口不是nil， 是带有type，而data==nil
*/
func main() {
	if live() == nil {
		fmt.Println("AAA")
	} else {
		fmt.Println("BBB")
	}
}