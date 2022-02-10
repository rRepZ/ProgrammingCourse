package main

import "fmt"

func main() {
	ff := inc()
	ff()
	ff()
	ff()

	fff := inc()
	fff()
	fff()
}

func inc() func() {
	var i int
	f := func() {
		i++
		fmt.Println(i)
	}

	return f
}
