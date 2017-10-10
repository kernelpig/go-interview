package main

import "fmt"

type student struct {
	Name string
	Age int
}

func pase_student_error(){
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "wang", Age: 1},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// 正确方式1
func pase_student_1(){
	m := make(map[string]*student)
	stus := []*student{
		{Name: "zhou", Age: 24},
		{Name: "wang", Age: 1},
	}
	for _, stu := range stus {
		// 不要对for中的变量取地址
		m[stu.Name] = stu
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// 正确方式2
func pase_student_2(){
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "wang", Age: 1},
	}
	for _, stu := range stus {
		// 不要对for中的变量取地址
		stu1 := stu
		m[stu.Name] = &stu1
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main(){
	pase_student_error()
}