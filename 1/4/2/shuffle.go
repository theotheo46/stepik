package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// shuffle перемешивает элементы nums in-place.
func shuffle(nums []int) {
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// init - это специальная функция, которую Go вызывает до main()
// обычно используется для инициализации начального состояния приложения
func init() {
	// Функция rand.Seed() инициализирует генератор случайных чисел
	// здесь мы используем константу 42, чтобы программу
	// можно было проверить тестами.
	//
	// В реальных задачах не используйте константы!
	// Используйте, например, время в наносекундах:
	// rand.Seed(time.Now().UnixNano())
	//
	// А лучше создайте новый генератор случайных чисел
	// и используйте его вместо глобального:
	// rnd := rand.New(rand.NewSource(42))
	rand.Seed(42)
}

func main() {
	nums := readInput()
	shuffle(nums)
	fmt.Println(nums)
}

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
