package main

import (
	"fmt"
	"sync"
)

// начало решения

type Counter struct {
	lock sync.RWMutex
	m    map[string]int
}

func (c *Counter) Increment(str string) {
	c.lock.Lock()
	c.m[str] = c.m[str] + 1
	c.lock.Unlock()
}

func (c *Counter) Value(str string) int {
	c.lock.RLock()
	ret := c.m[str]
	c.lock.RUnlock()
	return ret
}

func (c *Counter) Range(fn func(key string, val int)) {
	c.lock.RLock()
	for k, v := range c.m {
		fn(k, v)
	}
	c.lock.RUnlock()
}

func NewCounter() *Counter {
	c := Counter{m: make(map[string]int)}
	return &c
}

// конец решения

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(3)

	increment := func(key string, val int) {
		defer wg.Done()
		for ; val > 0; val-- {
			counter.Increment(key)
		}
	}

	go increment("one", 100)
	go increment("two", 200)
	go increment("three", 300)

	wg.Wait()

	fmt.Println("two:", counter.Value("two"))

	fmt.Print("{ ")
	counter.Range(func(key string, val int) {
		fmt.Printf("%s:%d ", key, val)
	})
	fmt.Println("}")
}
