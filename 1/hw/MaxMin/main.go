package main

import (
	"fmt"
	"strings"
)

func main() {

	StringExMin()

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
