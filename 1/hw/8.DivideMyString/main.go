package main

import (
	"fmt"
	"strings"
)

func main() {
	// i := "Ivan"
	// g := "Golang"

	// i += g // выделяется память, создаётся новый массив размером 10, в него копируются все элементы
	// // самое главное, это происходит неявно

	// love := i + g
	// fmt.Println(love)

	// var n, d, lenStr int
	// fmt.Scanf("%d\n", &n)
	// fmt.Scanf("%d\n", &d)
	// numStr := make([]string, n)
	// scanStr := bufio.NewScanner(os.Stdin)
	// scanStr.Scan()

	// numStr = strings.Split(scanStr.Text(), " ")

	// divideIndex := 0
	// result := make([]string, d)
	// lenStr = n
	// if n%d != 0 {
	// 	lenStr = d
	// }

	// // использовать стрингбилдер вместо конкатенации
	// for i := 0; i < d; i++ {
	// 	//fmt.Println("result n/d", n/d)
	// 	for ; divideIndex < lenStr/d; divideIndex++ {
	// 		//fmt.Println("res ", result)
	// 		result[i] += numStr[divideIndex]
	// 		fmt.Println("res ", result)
	// 	}

	// 	numStr = numStr[divideIndex:]
	// 	fmt.Println("div ", numStr)
	// 	divideIndex = 0
	// }
	// if n%d != 0 {
	// 	for i := d - 1; i < d; i++ {
	// 		for ; divideIndex < len(numStr); divideIndex++ {
	// 			//fmt.Println("res ", result)
	// 			result[i] += numStr[divideIndex]
	// 			fmt.Println("res ", result)
	// 		}
	// 	}
	// }

	// // for _, r := range result {
	// // 	strings.Join(r, " ")
	// // }

	phoneByName := make(map[string]string)

	// viktorsPhone := phoneByName["viktor"]
	phoneByName["viktor"] = ""

	viktorsPhone, ok := phoneByName["viktor"]
	fmt.Printf("is viktor's phone is known: %v, [%s]\n", ok, viktorsPhone)

	// if в одну строчку
	if v, ok := phoneByName["Fyodor"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}

	delete(phoneByName, "viktor")
	viktorsPhone, ok = phoneByName["viktor"]
	fmt.Printf("is viktor's phone is known after delete: %v, [%s]\n", ok, viktorsPhone)

	fmt.Printf("viktors phohe model is [%s]\n", viktorsPhone)
	// set("ivan", "samsung")
	phoneByName["ivan"] = "samsung"

	// get("ivan")
	// phone := phoneByName["ivan"]

}

func ConcatString() string {
	res := ""
	for i := 0; i < 10000; i++ {
		res += "6"
	}

	return res
}

func ConcatStringBuilder() string {
	var sb strings.Builder

	for i := 0; i < 10000; i++ {
		sb.WriteString("6")
	}

	return sb.String()
}

func ConcatStringBuilderAllocateMemory() string {
	var sb strings.Builder
	sb.Grow(10000)
	for i := 0; i < 10000; i++ {
		sb.WriteString("6")
	}

	return sb.String()
}
