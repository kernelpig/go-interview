package main

import "fmt"

/*
结果：./19.go:9:2: cannot use nil as type string in return argument
*/
func GetValue19(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return nil, false
	// OK
	// return "", false
}

func main(){
	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}
	v, err := GetValue19(intmap, 3)
	fmt.Println(v, err)
}
