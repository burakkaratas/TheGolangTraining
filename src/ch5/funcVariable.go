package main

import "fmt"

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

func main() {

	f1 := square
	f2 := negative
	f3 := product

	fmt.Println(f1(1))
	fmt.Println(f2(2))
	fmt.Println(f3(3,4))
}
