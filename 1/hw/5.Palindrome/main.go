package main

import (
	"fmt"

	//"strconv"
	"strings"
)

func main() {

	IsPalindrome()

}

func IsPalindrome() {
	// var str string
	// fmt.Scanf("%s", &str)
	str := "好ada好"
	words := strings.Split(str, "")

	n := len(words)
	var isPal bool
	for i := 0; i < n; i++ {
		if words[i] == words[n-1] {
			isPal = true
			n = n - 1
		} else {
			isPal = false
			break
		}
	}
	fmt.Println("Является ли строка палиндромом -- ", isPal)
}
