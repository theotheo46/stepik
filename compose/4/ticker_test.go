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
				at := time.Now()
				fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
			}
			cancel := schedule(50*time.Millisecond, work)
			cancel()
			defer cancel()
			time.Sleep(260 * time.Millisecond)
			cancel()
			time.Sleep(260 * time.Millisecond)
		})
	}
}
