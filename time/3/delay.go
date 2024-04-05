package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

func delay(dur time.Duration, fn func()) func() {
	cancelchan := make(chan struct{}, 1)
	timer := time.NewTimer(dur)
	isFirstTime := true

	cancel := func() {
		if isFirstTime {
			cancelchan <- struct{}{}
			isFirstTime = false
		}
	}

	go func() {
		select {
		case <-timer.C:
			fn()
			return
		case <-cancelchan:
			return
		}

	}()

	return cancel
}

// конец решения

func main() {
	rand.Seed(time.Now().Unix())

	work := func() {
		time.Sleep(5 * time.Millisecond)
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)
	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(100 * time.Millisecond)
}
