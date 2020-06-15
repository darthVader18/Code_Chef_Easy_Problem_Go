package main

import "fmt"

func main() {
	var t, k int
	fmt.Scan(&t)
	for index := 0; index < t; index++ {
		var n int
		fmt.Scan(&n)
		playlist := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&playlist[i])
			}
			fmt.Scan(&k)
		below := 0
		for i := 0; i < n; i++ {
			if playlist[i] <= playlist[k-1] {
				below ++
				}
		}
		fmt.Println(below)
	}
}

