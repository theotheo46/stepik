// начало решения

// исправьте count() и take(),
// чтобы main() могла их отменить

// count отправляет в канал числа от start до бесконечности
func count(start int) <-chan int {
	out := make(chan int)
	go func() {
		for i := start; ; i++ {
			out <- i
		}
	}()
	return out
}

// take выбирает первые n чисел из in и отправляет в выходной канал
func take(in <-chan int, n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- <-in
		}
		close(out)
	}()
	return out
}

// конец решения
