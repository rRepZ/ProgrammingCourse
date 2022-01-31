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

	rightSlice := make([]int, 0, len(saveSlice))
	leftSlice := make([]int, 0, len(saveSlice))
	glob_j := 0
	for i := 0; i < len(saveSlice); i++ {
		state := UNKNOWN
		intSlice := StringToInt(strings.Split(saveSlice[i], " "))
		fmt.Println(len(intSlice))
		if glob_j+1 >= len(intSlice) {
			leftSlice = append(leftSlice, intSlice...)

		} else if intSlice[glob_j] > intSlice[glob_j+1] {

			for intSlice[glob_j] > intSlice[glob_j+1] {
				if glob_j+2 < len(intSlice) {
					glob_j++

				} else {
					glob_j++
					break
				}
			}
			leftSlice = append(leftSlice, intSlice[:glob_j+1]...)
			rightSlice = append(rightSlice, intSlice[glob_j+1:]...)
			glob_j = 0
		} else if intSlice[glob_j] < intSlice[glob_j+1] {
			for intSlice[glob_j] <= intSlice[glob_j+1] {

				if glob_j+2 < len(intSlice) {
					glob_j++
				} else {
					glob_j++
					break
				}
				//glob_j++
			}
			leftSlice = append(leftSlice, intSlice[:glob_j+1]...)
			fmt.Println(leftSlice)
			rightSlice = append(rightSlice, intSlice[glob_j+1:]...)
			glob_j = 0
		} else if intSlice[glob_j] == intSlice[glob_j+1] {
			for intSlice[glob_j]-intSlice[glob_j+1] <= 0 {
				if glob_j+2 < len(intSlice) {
					glob_j++
				} else {
					glob_j++
					break
				}

				leftSlice = append(leftSlice, intSlice[:glob_j+1]...)
				fmt.Println(leftSlice)
				rightSlice = append(rightSlice, intSlice[glob_j+1:]...)
				glob_j = 0
				//glob_j++
			}
		}

	}
	//fmt.Println(intSlice)
	//fmt.Println(leftSlice)
	//fmt.Println(rightSlice)
	fmt.Println(append(rightSlice, leftSlice...))

}

func StringToInt(strTmp []string) []int {
	sliceOfInt := make([]int, 0, len(strTmp))
	var str string
	var i int
	for _, str = range strTmp {
		i, _ = strconv.Atoi(str)
		sliceOfInt = append(sliceOfInt, i)
	}
	return sliceOfInt
}
