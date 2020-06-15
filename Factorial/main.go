package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for index := 0; index < t; index++ {
		var n, sum int
		fmt.Scan(&n)
		for n >= 5 {
			sum = sum + n/5
			n = n / 5
		}
		fmt.Println(sum)
	}
}
