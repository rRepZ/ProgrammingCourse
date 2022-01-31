package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DESC = iota - 1
	UNKNOWN
	ASC
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	scanner := bufio.NewScanner(os.Stdin)

	rights := make([]int, 0)
	lefts := make([]int, 0)

	for i := 0; i < n; i++ {
		scanner.Scan()

		numsStr := strings.Split(scanner.Text(), " ")
		nums := make([]int, 0, len(numsStr))

		for _, ns := range numsStr {
			n, _ := strconv.Atoi(ns)
			nums = append(nums, n)
		}

		index := 0
		state := UNKNOWN

	LOOP:
		// for i := 0; i < len(nums) - 2; i++
		// желательно до len(nums) - 1
		for i := range nums {
			if i+2 > len(nums) {
				index = i + 1
				break
			}

			switch {
			case nums[i] > nums[i+1]:
				// избавиться от условия, идти в цикле до len(nums) - 1
				// в лучшем случае общая логика должна содержать именно условие что делать, если посл-ть низходящая но на данном элементе наоборот
				if i+2 < len(nums) {
					if state == DESC {
						index = i + 1
						break LOOP
					}
				} else if state == ASC {
					index = i + 2
					break LOOP
				} else {
					index = i + 1
					break LOOP
				}
				state = ASC
			case nums[i] < nums[i+1]:
				state = DESC
				if i+2 < len(nums) {
					if state == ASC {
						index = i + 1
						break LOOP
					}
				} else if state == DESC {
					index = i + 2
					break LOOP
				} else {
					index = i + 1
					break LOOP
				}
			}
		}

		// краевые случаи рассмотреть здесь
		lefts = append(lefts, nums[:index]...)
		rights = append(rights, nums[index:]...)
	}

	result := append(rights, lefts...)
	for _, i := range result {
		fmt.Printf("%d ", i)
	}
}
