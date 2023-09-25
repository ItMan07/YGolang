package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	D := math.Pow(b, 2) - 4.0*a*c
	if D > 0 {
		x1 := (-b - math.Pow(D, 0.5)) / (2 * a)
		x2 := (-b + math.Pow(D, 0.5)) / (2 * a)
		fmt.Println(x1, x2)
	} else if D == 0 {
		x1 := -b / (2 * a)
		fmt.Println(x1)
	} else {
		fmt.Println(0, 0)
	}
}

func SumDigitsRecursive(n int) int {
	if n < 10 {
		return n
	}

	// Рекурсивно вызываем функцию для суммы цифр числа без последней цифры,
	// а затем добавляем последнюю цифру к этой сумме.
	lastDigit := n % 10
	remainingDigits := n / 10
	return lastDigit + SumDigitsRecursive(remainingDigits)
}

func IsPowerOfTwoRecursive(N int) {
	for i := 1; i <= N; i *= 2 {
		if i == N {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func progress(n float32) float32 {
	sum := n * (1 + n) / 2
	return sum
}

func discriminant(a, b, c float64) float64 {
	// ax^2 + bx + c
	D := math.Pow(b, 2) - 4.0*a*c
	return D
}

func fx(x float32) (float32, float32) {
	y := x * 1.555
	return x, y
}

func main() {
	n := 0
	if fmt.Scan(&n); n >= 0 {
		sum := 0
		for i := 1; i <= n; i++ {
			if (i%3 != 0) && (i%5 != 0) {
				sum += i
			}
		}
		fmt.Println(sum)
	} else {
		fmt.Println("Некорректный ввод")

		fmt.Println(discriminant(2, 10, 1))
		fmt.Println(fx(10))
		fmt.Println(progress(8))
		fmt.Println(factorial(5))
		fmt.Println(fib(10))
	}
}
