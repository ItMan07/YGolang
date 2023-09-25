package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	num := 0
	fmt.Scan(&num)

	// Найти ближайшее большее число Фибоначчи
	n := 0
	for {
		if fibonacci(n) >= num {
			break
		}
		n++
	}

	// Вывести следующие 10 чисел Фибоначчи
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(n))
		n++
	}
}
