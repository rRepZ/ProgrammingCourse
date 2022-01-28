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
	fmt.Scanf("%d", &n)
	sliceOfStr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &baseStr)
		sliceOfStr[i] = baseStr
	}
	SplitMyString(sliceOfStr)
}

func SplitMyString(getSlice []string) {
	var capOfSlice int
	endSlice := make([]string, 0)
	endSliceL := make([]string, 0)
	capOfSlice = cap(getSlice)
	for i := 0; i < capOfSlice; i++ {
		str := getSlice[i]
		sliceOfStr := strings.Split(str, "")
		sizeCap := cap(sliceOfStr)

		sliceRight := make([]string, sizeCap)
		sliceLeft := make([]string, sizeCap)
		if (sizeCap > 1) && (sliceOfStr[0] < sliceOfStr[1]) {
			for i := 0; i < sizeCap; i++ {
				if (i+1 < sizeCap) && (sliceOfStr[i] < sliceOfStr[i+1]) {
					sliceLeft[i] = sliceOfStr[i]
				} else {
					sliceRight[i] = sliceOfStr[i]
				}
			}
		} else {
			for i := 0; i < sizeCap; i++ {
				if (i+1 < sizeCap) && (sliceOfStr[i] > sliceOfStr[i+1]) {
					sliceLeft[i] = sliceOfStr[i]
				} else {
					sliceRight[i] = sliceOfStr[i]
				}
			}
		}
		endSlice = append(endSlice, sliceRight...)
		endSliceL = append(endSliceL, sliceLeft...)
	}
	endSlice = append(endSlice, endSliceL...)
	fmt.Println(endSlice)

}
