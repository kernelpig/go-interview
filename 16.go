package main

import "fmt"

/*
结果：./16.go:11:13: cannot use s2 (type []int) as type int in append
*/
func main(){
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	// s1 = append(s1, s2...)
	fmt.Println(s1)
}
