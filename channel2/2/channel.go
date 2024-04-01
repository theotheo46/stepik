package main

import (
	"fmt"
	"time"
)

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// начало решения
	var result []any
	done := make(chan any, len(funcs))
	for _, f := range funcs {
		fmt.Println(f())
		done := make(chan any, 1)
		done_arr = append(done_arr, done)
		go func(f func() any) {
			done <- f()
		}(f)
	}

	for _, ch := range done_arr {
		result = append(result, <-ch)
	}

	return result
	// выполните все переданные функции,
	// соберите результаты в срез
	// и верните его

	// конец решения
}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(1), squared(10), squared(1)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
