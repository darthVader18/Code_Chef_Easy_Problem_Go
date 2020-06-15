package main

import "fmt"

func main() {
	var A, B, diff int
	fmt.Scan(&A)
	fmt.Scan(&B)
	diff = A - B
	if diff % 10 != 9 {
		diff = diff + 1
	} else {
		diff = diff - 1
	}
	fmt.Println(diff)
}
