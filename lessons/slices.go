package lessons

import (
	"fmt"
	"sort"
)

func example() {
	var index int
	fmt.Print("Введите индекс: ")
	fmt.Scanln(&index)
	numbers := []int{1, 2, 3, 4, 5}

	if index <= len(numbers)-1 {
		fmt.Println(numbers[index])
	} else {
		fmt.Println("Index out of range")
	}
	numbersNew := append(numbers, 10)
	fmt.Println(numbersNew)
	println(numbers)
}

func finder() {
	nums := []int{1, 2, 3, 4, 5}
	target := 3
	found := false

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Number", target, "is found :)")
	} else {
		fmt.Println("Number not found :(")
	}
}

func strings() {
	firstName := "Андрей"
	lastName := "Смирнов"
	fullName := firstName + " " + lastName
	fmt.Println(fullName) // Output: Андрей Смирнов
	fmt.Println(len(fullName))
}

/*
Функция len возвращает длину слайса, то есть количество элементов,
которые в данный момент доступны для использования.

Функция cap возвращает ёмкость слайса, то есть количество элементов,
которые могут быть сохранены в слайсе до того, как он должен быть перераспределён.

Функция append используется для добавления новых элементов в слайс.

Функция copy позволяет копировать элементы из одного слайса в другой.
*/

func ReverseSlice(slice []int) []int {
	reversed := make([]int, len(slice))

	for i := 0; i < len(slice); i++ {
		reversed[i] = slice[len(slice)-i-1]
	}

	return reversed
}

func FindMinMaxInSlice(slice []int) (int, int) {
	//mini := 0
	//maxi := 0
	//
	//slices.Sort(slice)
	//if len(slice) > 0 {
	//	mini = slice[0]
	//	maxi = slice[len(slice)-1]
	//}

	if len(slice) > 0 {
		mini := 999999
		maxi := 0
		for i := 0; i < len(slice); i++ {
			if slice[i] > maxi {
				maxi = slice[i]
			}
			if slice[i] < mini {
				mini = slice[i]
			}
		}
		return mini, maxi
	} else {
		return 0, 0
	}
}

func SumOfSlice(slice []int) int {
	summ := 0
	for i := 0; i < len(slice); i++ {
		summ += slice[i]
	}
	return summ
}

func IntersectionOfSlices(slice1, slice2 []int) []int {
	out := []int{}
	for i := 0; i < len(slice1); i++ {
		for j := 0; j < len(slice2); j++ {
			if slice1[i] == slice2[j] {
				out = append(out, slice1[i])
			}
		}
	}
	return out
}

func SortSlice(slice []int) []int {
	sort.Ints(slice)
	return slice
}

func f9() {
	//nums := []int{1, 2, 6, 4, 5}
	//nums := []int{1, 2, 6, 4, 5}
	nums2 := []int{-4, -2, -10}
	//nums2 := []int{1, 2}
	//fmt.Println(IntersectionOfSlices(nums, nums2))
	//fmt.Println(SumOfSlice([]int{1, 2, 5, 4, 3}))
	fmt.Println(SortSlice(nums2))
}

func f10() {
	finder()

	var nums = [5]int{1, 2, 3, 4, 5}
	dynamicNums := []int{1, 2, 3}
	var fruits = []string{"apple", "banana", "orange"}

	fmt.Println(len(fruits), cap(fruits))
	fruits = append(fruits, "olive")
	fmt.Println(nums[0], dynamicNums[1])
	fmt.Println(fruits)
}
