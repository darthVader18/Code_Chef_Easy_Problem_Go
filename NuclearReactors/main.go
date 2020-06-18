package main

import (
"fmt"
)

func main() {

	var A, N, K int
	fmt.Scan(&A, &N, &K)
	output := make([]int, K)
	i := 0
	for ; A > 0 && i < K; {
		output[i] = A % (N + 1)
		A = A /(N + 1)
		i++
	}
	for i = 0; i < K; i++{
		fmt.Printf("%d ", output[i])
	}
	fmt.Println()
}