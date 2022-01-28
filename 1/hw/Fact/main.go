package main

import (
	"fmt"
)

func main() {

	FactRecResult()
	Factorial()

}

func FactRecResult() {
	var n int
	fmt.Scanf("%d", &n)
	if n == 0 {
		fmt.Println(1)
		return
	}
	n = FactorialRecurs(n)
	if n != -1 {
		fmt.Println("Факториал равен", n)
	} else {
		fmt.Println("Факториала не существует")
	}
}

func FactorialRecurs(n int) int {
	return n * FactorialRecurs(n-1)
}

func Factorial() {
	var n, f int
	fmt.Scanf("%d", &n)
	f = n
	if n == 0 {
		f = 1
	} else if n > 0 {

		for i := 1; i < n; i++ {
			f = f * i
		}
	}
	if n >= 0 {
		fmt.Println("факториал равен ", f)
	} else {
		fmt.Println("Факториала не существует")
	}
}
