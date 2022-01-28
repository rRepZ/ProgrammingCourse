package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	if n == 0 {
		fmt.Println(1)
		return
	}

	var f int
	for i := 1; i < n; i++ {
		f = f * i
	}

	fmt.Println(f)
}
