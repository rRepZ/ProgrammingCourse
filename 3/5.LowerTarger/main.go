package main

import (
	"fmt"
)

//leetcode
func lowerThenTargetMass(s []int, target int) []int {
	resultMap := make(map[int]int)
	resultNums := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		if i1, ok := resultMap[target-s[i]]; !ok {
			resultMap[s[i]] = i
		} else {
			fmt.Printf("%d:%d \n", i1, i)
			fmt.Println(resultMap)
			return
		}

	}

}

func main() {
	nums := []int{2, 3, 4, 83, 10, 3, 8, -3, -87, 999, 83, 999, -1, -1, -19, -1, 999}
	target := 85
	twoSum(nums, target)

}
