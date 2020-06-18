package main

import "fmt"

func main() {
	var t, n int
	var s string
	fmt.Scan(&t)
	for index := 0; index < t; index++ {
		fmt.Scan(&n) //the length of string
		fmt.Scan(&s)
		count := 0
		for i := 0; i < n; i++ {
			if s[i] == '1' {
				count ++
			}
		}
		fmt.Println((count * (count + 1)) / 2)
	}
}

