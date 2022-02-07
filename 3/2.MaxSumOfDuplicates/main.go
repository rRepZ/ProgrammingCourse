package main

import "fmt"

func MaxSum(thisSlice []int) {
	resultMap := make(map[int]int)
	for i := 0; i < len(thisSlice); i++ {
		if resultMap[thisSlice[i]] >= thisSlice[i] {
			resultMap[thisSlice[i]] += thisSlice[i]
		} else {
			resultMap[thisSlice[i]] = thisSlice[i]
		}
	}
	fmt.Println(resultMap)
}

func main() {
	mass := []int{1, 2, 3, 4, 5, 2, 6, 7, 8, 8}
	MaxSum(mass)
	n := "b"
	fmt.Println([]rune(n)[0] - []rune("a")[0])

}
