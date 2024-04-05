package main

import (
	"fmt"
	"sync"
)

// начало решения

type Counter struct {
	m map[string]int
	l sync.Mutex
}

func (c *Counter) Increment(str string) {
	c.l.Lock()
	c.m[str] = c.m[str] + 1
	c.l.Unlock()
}

func (c *Counter) Value(str string) int {
	c.l.Lock()
	ret := c.m[str]
	c.l.Unlock()
	return ret
}

func (c *Counter) Range(fn func(key string, val int)) {
	c.l.Lock()
	for k, v := range c.m {
		fn(k, v)
	}
	c.l.Unlock()
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
