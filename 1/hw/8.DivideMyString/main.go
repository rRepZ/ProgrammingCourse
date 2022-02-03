package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, d, lenStr int
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%d\n", &d)
	numStr := make([]string, n)
	scanStr := bufio.NewScanner(os.Stdin)
	scanStr.Scan()

	numStr = strings.Split(scanStr.Text(), " ")

	numStr = numStr[:n]

	divideIndex := 0
	result := make([]string, d)
	lenStr = n
	if n%d != 0 {
		lenStr = d
	}

	for i := 0; i < d; i++ {
		//fmt.Println("result n/d", n/d)
		for ; divideIndex < lenStr/d; divideIndex++ {
			//fmt.Println("res ", result)
			result[i] += " " + numStr[divideIndex]
			fmt.Println("res ", result)
		}
		numStr = numStr[divideIndex:]
		fmt.Println("div ", numStr)
		divideIndex = 0
	}
	if n%d != 0 {
		for i := d - 1; i < d; i++ {
			for ; divideIndex < len(numStr); divideIndex++ {
				//fmt.Println("res ", result)
				result[i] += " " + numStr[divideIndex]
				fmt.Println("res ", result)
			}
		}
	}

	for i := d - 1; i >= 0; i-- {
		fmt.Println(result[i])
	}

}
