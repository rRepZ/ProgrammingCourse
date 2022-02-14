package main

import (
	"fmt"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
}

func BenchmarkField(b *testing.B) {
	field := make([][]int, 10000)

	for i := 0; i < 10000; i++ {
		//field = append(field, []string{}) //выделяем память по i
		for j := 0; j < 10000; j++ {
			field[i] = append(field[i], 999)
		}
	}

}

func BenchmarkField2(b *testing.B) {
	field := make([][]int, 0, 10000)

	for i := 0; i < 10000; i++ {
		field = append(field, []int{}) //выделяем память по i
		for j := 0; j < 10000; j++ {
			field[i] = append(field[i], 999)
		}
	}

}

func BenchmarkField3(b *testing.B) {
	a := make([][]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		arr := make([]int, 0, 10000)
		a = append(a, arr)
		for j := 0; j < 10000; j++ {
			a[i] = append(a[i], 999)
		}
	}

}

func BenchmarkField4(b *testing.B) {
	//field := make([][]int, 0, 100)
	var field [][]int
	for i := 0; i < 10000; i++ {
		field = append(field, []int{}) //выделяем память по i
		for j := 0; j < 10000; j++ {
			field[i] = append(field[i], 999)
		}
	}

}

func BenchmarkField5(b *testing.B) {
	field := make([][]int, 10000)
	//var field [][]int
	for i := 0; i < 10000; i++ {
		field = append(field, []int{}) //выделяем память по i
		for j := 0; j < 10000; j++ {
			field[i] = append(field[i], 999)
		}
	}

}
