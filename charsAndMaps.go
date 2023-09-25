package main

import "fmt"

func main() {
	name := "異體字"
	firstLetter := name[0]
	fmt.Println(firstLetter)         // Вывод: 231 (код символа "А" в таблице Unicode)
	fmt.Println(string(firstLetter)) // Вывод: ç
}
