package main

func main() {
	a() // stack_push(pointer_main)

}

func a() {
	a() // stack_push(pointer_a)
}

func b() {
	// stack_pop
}
