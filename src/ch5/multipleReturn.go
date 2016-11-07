package main

import "fmt"

func sumWithParam(param1, param2 int) (int, int, int) {
	return param1, param2, param1 + param2
}

func main() {
	x, y, sum := sumWithParam(2, 3)
	fmt.Printf(" %d + %d = %d ", x, y, sum)
}
