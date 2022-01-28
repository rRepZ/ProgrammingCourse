package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"math"
	//"strconv"
	"strings"
)

func main() {
	//Triangle()
	//Square()
	//Factorial()
	//FactRecResult()
	//IsPalindrome()
	//RevString()
	//StringExMin()
	//DivideString()
}

func DivideString() {
	var numberOfSymb, divStr, f int
	//var baseStr string
	var endStr [8][8]string
	//endStr := make([][]string, 20)
	fmt.Scanf("%d", &numberOfSymb)
	fmt.Scanf("%d", &divStr)
	//fmt.Scanf("%s", &baseStr)
	baseStr := bufio.NewScanner(os.Stdin)
	for baseStr.Scan() {
		fmt.Println(baseStr.Text())
		break
	}
	sliceOfStr := strings.Split(baseStr.Text(), " ")

	fmt.Println(sliceOfStr)
	fmt.Println(endStr)
	for i := 0; i < numberOfSymb; {
		f = i
		for j := 0; j < divStr; j++ {
			endStr[f][j] = sliceOfStr[i]
			if i+1 < numberOfSymb {
				fmt.Println("i = ", i)
				i++
			} else {
				i++
				break
			}
			//i++

		}
	}
	fmt.Println(endStr)
}

func StringExMin() {
	var n int
	var baseStr string
	//var sliceRight, sliceLeft []string
	//var sliceOfStr []string
	fmt.Println("Сколько строк вы хотите ввести? ")
	fmt.Scanf("%d", &n)
	sliceOfStr := make([]string, n)

	fmt.Println("Введите строки: ")
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &baseStr)
		sliceOfStr[i] = baseStr
	}

	fmt.Println("Срез строк: ", sliceOfStr)
	SplitMyString(sliceOfStr)
}

func SplitMyString(getSlice []string) {
	var capOfSlice, sizeCap int
	var str string
	endSlice := make([]string, 0)
	endSliceL := make([]string, 0)
	capOfSlice = cap(getSlice)
	//slice_of_slices := make([][]int , 3)
	for i := 0; i < capOfSlice; i++ {
		str = getSlice[i]
		sliceOfStr := strings.Split(str, "")
		sizeCap = cap(sliceOfStr)

		sliceRight := make([]string, sizeCap)
		sliceLeft := make([]string, sizeCap)
		fmt.Println("Size of Slice: ", sizeCap)
		if (sizeCap > 1) && (sliceOfStr[0] < sliceOfStr[1]) {
			for i := 0; i < sizeCap; i++ {
				if (i+1 < sizeCap) && (sliceOfStr[i] < sliceOfStr[i+1]) {
					sliceLeft[i] = sliceOfStr[i]
				} else {
					sliceRight[i] = sliceOfStr[i]
				}
				//fmt.Println("i сейчас равна ", i)
			}
		} else {
			for i := 0; i < sizeCap; i++ {
				if (i+1 < sizeCap) && (sliceOfStr[i] > sliceOfStr[i+1]) {
					sliceLeft[i] = sliceOfStr[i]
				} else {
					sliceRight[i] = sliceOfStr[i]
				}
				//fmt.Println("i сейчас равна ", i)
			}
		}
		fmt.Println("левая: ", sliceLeft)
		fmt.Println("правая: ", sliceRight)
		endSlice = append(endSlice, sliceRight...)
		endSliceL = append(endSliceL, sliceLeft...)
		//sliceLeft = append(sliceRight, sliceLeft...)
		//fmt.Println("правая + левая: ", sliceLeft)
	}
	endSlice = append(endSlice, endSliceL...)
	fmt.Println("правая + левая: ", endSlice)

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

func IsPalindrome() {
	var str string
	var isPal bool
	fmt.Scanf("%s", &str)
	words := strings.Split(str, "")
	n := len(words)
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

func FactRecResult() {
	var n int
	fmt.Scanf("%d", &n)
	n = FactorialRecurs(n)
	if n != -1 {
		fmt.Println("Факториал равен", n)
	} else {
		fmt.Println("Факториала не существует")
	}
}

func FactorialRecurs(n int) int {
	if n == 0 {
		return 1
	} else if n > 0 {
		return n * FactorialRecurs(n-1)
	} else {
		return -1
	}

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

func Triangle() {
	var sq float64
	slice := make([]string, 3)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(s.Text())
		break
	}
	numbers := strings.Split(s.Text(), " ")
	fmt.Println(numbers)
	fmt.Println("Размер ", cap(numbers))
	//i, err := strconv.Atoi(numbers[1])
	if cap(numbers) == 1 {
		numbers = append(numbers, slice...)
		//numbers[1] = numbers[0]
		//numbers[2] = numbers[0]
		a, err := strconv.ParseFloat(numbers[0], 64)
		if err == nil {
			sq = (float64(a*a) * math.Pow(3, 0.5)) / 4
			println("Площадь треугольника равна: ", sq)
		}

	} else if cap(numbers) < 3 {
		a, err := strconv.ParseFloat(numbers[0], 64)
		b, err := strconv.ParseFloat(numbers[1], 64)
		fmt.Println("a равна ", a)
		fmt.Println("b равна ", b)
		fmt.Println("Err равна ", err)
		if err == nil {
			//c := math.Pow(float64(a*a+b*b),0.5)
			sq = a * b / 2
			fmt.Println("Площадь треугольника равна", sq)
		}
	} else {
		a, err := strconv.ParseFloat(numbers[0], 64)
		b, err := strconv.ParseFloat(numbers[1], 64)
		c, err := strconv.ParseFloat(numbers[2], 64)
		if (err == nil) && (a+b > c) && (b+c > a) && (a+c > b) {
			p := (a + b + c) / 2
			sq = math.Pow((p * (p - a) * (p - b) * (p - c)), 0.5)
			fmt.Println("Площадь треугольника равна", sq)
		} else {
			fmt.Println("Треугольника не существует")
		}
	}
	fmt.Println("num", numbers)

}

func Square() {

	var d, x1, x2, a, b, c float64 //ax^2+bx+c = 0;  d = b^2 - 4ac
	fmt.Scanf("%f %f %f", &a, &b, &c)
	fmt.Println("a = ", a, "b = ", b, "c = ", c)
	d = b*b - 4*a*c
	if d >= 0 {
		x1 = (-b + math.Pow(d, 0.5)) / (2 * a)
		x2 = (-b - math.Pow(d, 0.5)) / (2 * a)
		fmt.Printf("%f %f \n", x1, x2)
	} else {
		fmt.Println("дискриминант меньше нуля")
	}
}
