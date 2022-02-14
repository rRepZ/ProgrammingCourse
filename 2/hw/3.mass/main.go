package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	size := 10
	m := make([][]int, 0, size)
	for i := 0; i < size; i++ {
		m = append(m, []int{})
		for j := 0; j < size; j++ {
			m[i] = append(m[i], rand.Intn(10))
		}
	}
	for i := 0; i < size; i++ {

		fmt.Println(m[i])

	}
}
