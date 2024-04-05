package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestGenerate(t *testing.T) {
	defer goleak.VerifyNone(t)

	rand.Seed(time.Now().Unix())

	work := func() {
		fmt.Println("work done")
	}

	for i := 0; i < 20; i++ {

		cancel := delay(100*time.Millisecond, work)
		time.Sleep(10 * time.Millisecond)
		if rand.Float32() < 0.5 {
			cancel()
			fmt.Println("delayed function canceled")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
