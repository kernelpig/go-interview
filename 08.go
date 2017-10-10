package main

import (
	"sync"
	_ "strconv"
	_ "fmt"
	"strconv"
	"fmt"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

// ERROR
func (ua *UserAges) errGet(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

// OK
func (ua *UserAges) okGet(name string) int {
	ua.Lock()
	defer ua.Unlock()

	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
/*
结果： 可能出现fatal error: concurrent map read and map write
解答： map元素位置不固定， 写入导致元素重排， 并发读会出现问题
注意： 1. 故map需要读写都加锁；
*/
func main() {
	user := UserAges{}
	user.ages = make(map[string]int, 0)
	user.Add("1", 1)
	user.errGet("1")

	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func(user *UserAges, i int) {
			user.Add(strconv.Itoa(i), i)
			wg.Done()
		}(&user, i)
	}

	for i := 0; i < 10; i++ {
		go func(user *UserAges, i int) {
			age := user.errGet(strconv.Itoa(i))
			fmt.Println(age)
			wg.Done()
		}(&user, i)
	}
	wg.Wait()
}
