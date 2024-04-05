package main

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestGenerate(t *testing.T) {
	defer goleak.VerifyNone(t)
	for i := 0; i < 10; i++ {
		t.Run("test.name", func(t *testing.T) {
			work := func() {
				fmt.Print(".")
			}

			handle, cancel := withRateLimit(5, work)
			defer cancel()

			start := time.Now()
			const n = 10
			for i := 0; i < n; i++ {
				handle()
			}
			fmt.Println()
			fmt.Printf("%d queries took %v\n", n, time.Since(start))
		})
	}
}
