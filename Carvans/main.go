package main

import "fmt"

func main() {
	var t, n, n2, speed int
	fmt.Scan(&t)
	for index := 0; index < t; index++ {
		fmt.Scan(&n)
		cars := 0
		for j := 0; j < n; j++ {
			fmt.Scan(&n2)
			if j == 0 {
				speed = n2
				cars++
			} else if speed >= n2 {
				speed = n2
				cars++
			}
		}
		fmt.Println(cars)
	}
}