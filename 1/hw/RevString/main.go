package main

import (
	"fmt"
	"strings"
)

func main() {

	RevString()

}

func RevString() {
	var n, f int
	var baseStr string
	fmt.Println("Длина строки = ")
	fmt.Scanf("%d", &n)
	f = n
	fmt.Scanf("%s", &baseStr)
	strTmp := make([]string, n)
	strEnd := make([]string, n)
	letter := strings.Split(baseStr, "")
	copy(strTmp, letter)
	fmt.Println("введённая строка: ", letter)
	fmt.Println("строка заданного количества символов: ", strTmp)
	for i := 0; i < n; i++ {
		//fmt.Println("i ", strTmp[i])
		//fmt.Println("f-1 ", strTmp[f-1])
		strEnd[i] = strTmp[f-1]
		f--
	}
	fmt.Println("перевёрнутая строка: ", strEnd)
}
