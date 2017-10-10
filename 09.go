package main

import (
	"sync"
	"fmt"
)

type threadSafeSet struct {
	s []int
	sync.RWMutex
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

// OK
func (set *threadSafeSet) okIter() <-chan interface{} {
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()
		for _, e := range set.s {
			ch <- e
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

/*
1. chan分配缓冲区， 否则会出现fatal error: all goroutines are asleep - deadlock!
*/
func main(){
	t := threadSafeSet{}
	t.s = []int{1, 2, 3}
	fmt.Println(t.s)
	c := t.Iter()

	for v := range c {
		fmt.Println(v)
	}
}
