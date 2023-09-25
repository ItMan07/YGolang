package lessons

import (
	"fmt"
	"os"
)

func f1() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i*j > 20 {
				continue
				//break
			}
			fmt.Println(i * j)
		}
	}

	var input string
	for {
		fmt.Scan(&input)
		if input == "exit" {
			//break
			os.Exit(1)
			fmt.Println("exit098")
		}
	}
}

func f2() {
	num := 0

	if fmt.Scan(&num); num < 0 {
		fmt.Println("Некорректный ввод")
	} else if num >= 0 && num < 10 {
		fmt.Println("Число меньше 10")
	} else if num >= 10 && num < 100 {
		fmt.Println("Число меньше 100")
	} else if num >= 100 && num < 1000 {
		fmt.Println("Число меньше 1000")
	} else if num >= 1000 {
		fmt.Println("Число больше или равно 1000")
	}
}

func f3() {
	var num int
	fmt.Println("Input value")
	fmt.Scan(&num)

	if num >= 10 && num < 20 {
		fmt.Println("10 <= x < 20")

	} else if num >= 20 || num <= 30 {
		fmt.Println("20 <= x <= 30")

		var str string
		fmt.Scan(&str)

		switch str {
		case "a":
			fmt.Println("switch a")
		case "b":
			fmt.Println("switch b")
		default:
			fmt.Println("default value")
		}

		switch {
		case str == "c":
			fmt.Println("other switch case")
		}

	} else {
		fmt.Println("x - any value")
	}
}

func f4() {
	var rain bool
	fmt.Print("Идет дождь? ")
	fmt.Scan(&rain)
	if rain {
		fmt.Println("Возьми зонтик :)")
	} else if rain == false {
		fmt.Println("Не бери зонтик")
	} else {
		fmt.Println("Что за фигня")
	}
}

func f5() {
	a := 10
	b := 5
	sum := a + b
	fmt.Println("Сумма:", sum)
}

func f6() {
	var meters float64 = 0.0
	fmt.Scan(&meters)

	const exchangeRate = 1852
	seaMiles := meters / exchangeRate
	fmt.Println(seaMiles)
}

func f7() {
	var a, b int
	fmt.Scanln(&a, &b)
	sum := a + b
	fmt.Println("Сумма чисел:", sum)
}

func f8() {
	var num int
	fmt.Scan(&num)

	doubled := num * 2
	fmt.Println("Удвоенное число:", doubled)
}

func helloWorld() {
	fmt.Println("Hello, World!")
}
