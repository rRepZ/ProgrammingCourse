package main

import "fmt"

// todo написать функцию, которая принимает на вход массив целых чисел, а возвращает числа без дубликатов.

func main() {
	//fmt.Println(RemoveDuplicatesBrutForce([]int{1, 2, 2, 1, 1}))
	fmt.Println(RemoveDuplicatesMap([]int{1, 2, 2, 1, 1, 3, 3, 4}))
}

func RemoveDuplicatesBrutForce(input []int) []int { // сложность временнная О(n^2); по памяти - O(n)
	result := make([]int, 0, 5)

	for i := 0; i < len(input); i++ {
		j := 0
		for ; j < len(result); j++ {
			if result[j] == input[i] {
				break
			}

		}
		if j == len(result) {
			result = append(result, input[i])
		}
	}
	return result
}

func RemoveDuplicatesMap(input []int) []int { // сложность врменная -- O(n); по памяти -- O(2n)
	resultMap := make(map[int]int)
	//element := 1
	for i := 0; i < len(input); i++ {
		if resultMap[input[i]] >= 1 {
			resultMap[input[i]] = resultMap[input[i]] + 1
		} else {
			resultMap[input[i]] = 1
		}
	}
	fmt.Println(resultMap)
	result := make([]int, 0, len(resultMap))
	for k := range resultMap {
		result = append(result, k)
	}

	return result
}

/*
// c использованием сортировки; golang sort + написать бенчмарки
func RemoveDuplicatesSort(input []int) []int {

}
*/
// написать программу для вычисления чисел Фиббоначи рекурсивно и циклом + бенчмарки
