package main

import (
	"fmt"
	"time"
)

// начало решения

func schedule(dur time.Duration, fn func()) func() {
	canceled := make(chan struct{})
	ticker := time.NewTicker(dur)
	go func() {
		for canceled != nil {
			select {
			case <-ticker.C:
				fn()
			case <-canceled:
				return
			}
		}
	}()
	cancel := func() {
		ticker.Stop()
		if canceled != nil {
			close(canceled)
			canceled = nil
		}
	}

	return cancel
}

// конец решения

func main() {
	work := func() {
		at := time.Now()
		fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
	}

	cancel := schedule(50*time.Millisecond, work)
	cancel()
	defer cancel()

	// хватит на 5 тиков
	time.Sleep(260 * time.Millisecond)
	cancel()
	time.Sleep(260 * time.Millisecond)
}
