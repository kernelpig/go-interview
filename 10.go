package main

import "fmt"

type People10 interface {
	Speak(string) string
}

type Student struct {

}

func (stu *Student) Speak(think string) string {
	if think == "bitch" {
		return "You are a good student"
	} else {
		return "hi"
	}
}

func main() {
	// ok
	var peo People10 = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

	/*
	// error
	var peo People = Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
	*/
}
