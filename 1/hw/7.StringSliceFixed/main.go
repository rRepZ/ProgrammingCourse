package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NEGATIVE = -1
	POSITIVE = 1
	UNKNOWN  = 0
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	s := bufio.NewScanner(os.Stdin)

	lefts := make([]int, 0)
	rights := make([]int, 0)
	for i := 0; i < n; i++ {
		s.Scan()
		numsStr := strings.Split(s.Text(), " ")
		nums := make([]int, 0, len(numsStr))
		for _, ns := range numsStr {
			n, _ := strconv.Atoi(ns)
			nums = append(nums, n)
		}

		threshold := len(nums)
		asc, desc := false, false
	LOOP:
		for i := 0; i < len(nums)-1; i++ {
			switch {
			case nums[i+1] > nums[i]:
				if desc {
					threshold = i + 1
					break LOOP
				}
				asc = true
			case nums[i+1] < nums[i]:
				if asc {
					threshold = i + 1
					break LOOP
				}
				desc = true
			}
		}

		// краевые случаи
		lefts = append(lefts, nums[:threshold]...)
		rights = append(rights, nums[threshold:]...)
	}

	for _, r := range rights {
		fmt.Printf("%d ", r)
	}
	for _, l := range lefts {
		fmt.Printf("%d ", l)
	}
}
