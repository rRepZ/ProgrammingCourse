package main

import (
	"fmt"
)

func RemoveDuplicatesMap(input []string) string { // сложность врменная -- O(n); по памяти -- O(2n)
	resultMap := make(map[string]int)
	var str string
	//element := 1
	for i := 0; i < len(input); i++ {
		if resultMap[input[i]] >= 1 {
			resultMap[input[i]] = resultMap[input[i]] + 1
		} else {
			resultMap[input[i]] = 1
		}
		if resultMap[input[i]]*(100/len(input)) > 60 {
			if str != input[i] {
				str += input[i]
			}
		}
	}
	fmt.Println(resultMap)
	/*
		result := make([]string, 0, len(resultMap))
		for k := range resultMap {
			result = append(result, k)
		}
	*/

	return str
}

func main() {
	str := []string{"h", "h", "t", "h", "f", "h", "h", "h", "n", "h", "s", "a", "q", "h", "e", "h", "h", "b", "h", "h", "f", "d", "h", "h", "h", "a", "h", "h", "h", "h", "h", "h", "h", "h", "h", "h", "h", "h", "h", "h", "h", "h", "h", "h"}
	fmt.Println(RemoveDuplicatesMap(str))
}
