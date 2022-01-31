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

func StringToInt(strTmp []string) []int {
	sliceOfInt := make([]int, 0, len(strTmp))
	for _, str := range strTmp {
		i, _ := strconv.Atoi(str)
		sliceOfInt = append(sliceOfInt, i)
	}
	return sliceOfInt
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	saveSlice := make([]string, n)
	for i := 0; i < n; i++ {
		baseStr := bufio.NewScanner(os.Stdin)
		baseStr.Scan()
		for j := 0; j < len(strings.Split(baseStr.Text(), " ")); j++ {
			if j != 0 {
				saveSlice[i] = saveSlice[i] + " " + strings.Split(baseStr.Text(), " ")[j]
			} else {
				saveSlice[i] = saveSlice[i] + strings.Split(baseStr.Text(), " ")[j]
			}
		}

	}
	fmt.Println("save_slice: ", saveSlice)
	/*
		if len(saveSlice) == 1 {
			fmt.Println("зашли")
			fmt.Println("result: ", saveSlice)
			return
		}
	*/
	rightSlice := make([]int, 0, len(saveSlice))
	leftSlice := make([]int, 0, len(saveSlice))

	for _, elem := range saveSlice {
		intSlice := StringToInt(strings.Split(elem, " "))

		index := 0
		state := UNKNOWN
		for i := range intSlice {
			if i+2 > len(intSlice) {
				fmt.Println("зашли")
				index = i + 1
				break
			}
			if intSlice[i] > intSlice[i+1] {
				if i+2 < len(intSlice) {
					if state == UNKNOWN {
						state = POSITIVE
					}

					if state == NEGATIVE {
						index = i + 1
						break
					}
				} else if state == POSITIVE {
					index = i + 2
					break
				} else {
					fmt.Println("тут ")
					index = i + 1
					break
				}
			} else if intSlice[i] < intSlice[i+1] {
				if i+2 < len(intSlice) {
					if state == UNKNOWN {
						state = NEGATIVE
					}

					if state == POSITIVE {
						index = i + 1
						break
					}
				} else if state == NEGATIVE {
					index = i + 2
					break
				} else {
					fmt.Println(len(intSlice))
					fmt.Println("тут1 ")
					index = i + 1
					break
				}

			} /*else if intSlice[i] == intSlice[i+1] {
				if i+2 < len(intSlice) {
					if state == NEGATIVE {
						state = UNKNOWN
					}
					if state == POSITIVE {
						state = UNKNOWN
					}
					if state == UNKNOWN {
						index = i + 1
						break
					}
				} else {
					index = i + 2
					break
				}
			} */
		}

		leftSlice = append(leftSlice, intSlice[:index]...)
		rightSlice = append(rightSlice, intSlice[index:]...)
	}

	fmt.Println("left_slice: ", leftSlice)
	fmt.Println("right_slice: ", rightSlice)
	fmt.Println("result: ", append(rightSlice, leftSlice...))
}
