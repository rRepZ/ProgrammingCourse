package main

import "fmt"

func MaxNum(s []int) {
	resultMap := make(map[int]int)
	max := 0
	resultMap[0] = s[0]
	for i := 1; i < len(s); i++ {
		if resultMap[0] <= s[i] {
			if resultMap[0] == s[i] {
				resultMap[s[i]] += 1
			}
			resultMap[0] = s[i]
			max = s[i]
		}

	}

	fmt.Println(resultMap[max] + 1)
}

func main() {
	mass := []int{999, 2, 3, 4, 83, 10, 3, 8, -3, -87, 999, 83, 999, -1, -1, -19, -1, 999}
	MaxNum(mass)

}
