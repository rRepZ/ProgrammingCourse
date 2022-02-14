package main

import (
	"fmt"
)

func main() {
	TwoDSlice()
}

func TwoDSlice() {
	a := make([][]int, 0, 10)
	for i := 0; i < 10; i++ {
		arr := make([]int, 0, 10)
		a = append(a, arr)
		for j := 0; j < 10; j++ {
			a[i] = append(a[i], 999)
		}
	}
	for i := 0; i < 10; i++ {
		fmt.Println(a[i])
	}

}
