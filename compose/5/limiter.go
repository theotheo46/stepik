package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCanceled error = errors.New("canceled")

// начало решения

func withRateLimit(limit int, fn func()) (handle func() error, cancel func()) {
	canceled := make(chan struct{})
	limiter := time.NewTicker(time.Duration(1000/limit) * time.Millisecond)
	handle = func() error {
		if canceled != nil {
			select {
			case <-limiter.C:
				go fn()
			case <-canceled:
				return nil
			}
		} else {
			return ErrCanceled
		}
		return nil
	}
	cancel = func() {
		limiter.Stop()
		if canceled != nil {
			close(canceled)
			canceled = nil
		}
	}
	return handle, cancel
}

// конец решения

func main() {
	work := func() {
		fmt.Print(".")
	}

	handle, cancel := withRateLimit(5, work)
	defer cancel()

	start := time.Now()
	const n = 10
	for i := 0; i < n; i++ {
		handle()
		//fmt.Print(i)
	}
	fmt.Println()
	fmt.Printf("%d queries took %v\n", n, time.Since(start))
}
