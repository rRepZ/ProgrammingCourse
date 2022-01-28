package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	sides := strings.Split(s.Text(), " ")

	switch len(sides) {
	case 1:
		a, _ := strconv.Atoi(sides[0])
		sq := (float64(a*a) * math.Sqrt(3)) / 4
		println(sq)
		return
	case 2:
		a, _ := strconv.Atoi(sides[0])
		b, _ := strconv.Atoi(sides[1])
		fmt.Println(math.Sqrt(float64(a*a + b*b)))
		return
	case 3:
		a, _ := strconv.Atoi(sides[0])
		b, _ := strconv.Atoi(sides[1])
		c, _ := strconv.Atoi(sides[2])

		if (a+b > c) && (b+c > a) && (a+c > b) {
			fmt.Println(-1)
			return
		}

		p := (a + b + c) / 2
		fmt.Println(math.Sqrt(float64((p * (p - a) * (p - b) * (p - c)))))
	}
}
