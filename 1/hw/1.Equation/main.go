package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64 //ax^2+bx+c = 0;  d = b^2 - 4ac
	fmt.Scanf("%f %f %f", &a, &b, &c)
	d := b*b - 4*a*c

	if d < 0 {
		fmt.Println(-1)
		return
	}

	if a == 0 && b == 0 {
		fmt.Println("fail")
		return
	}

	if a == 0 {
		x := -c / b
		fmt.Println(x)
		return
	}

	x1 := (-b + math.Sqrt(d)) / (2 * a)
	x2 := (-b - math.Sqrt(d)) / (2 * a)

	fmt.Printf("%f %f \n", x1, x2)
}
