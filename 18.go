package main

import "fmt"

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty")
		return
	}
	fmt.Println("non-empty")
}

func main(){
	var x *int = nil
	Foo(x)
}
