package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = 2
		//		pc[i] = pc[i / 2] + byte(i & 1)
	}
}

func main() {
	for i := range pc {
		fmt.Println(pc[i])
	}
}
