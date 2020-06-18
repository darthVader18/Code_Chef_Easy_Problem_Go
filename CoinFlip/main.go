package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for index := 0; index < T; index++ {
		var G int //no. of games played by an elephant
		fmt.Scan(&G)
		for index2 := 0; index2 < G; index2++ {
			var I, N, Q, ans int //I denotes initial state of the coin, N denotes the no. coins and rounds
								 //if Q=1 (head), Q=2 (Tail)
								 //if I=1 (head), I=2 (Tail)
			fmt.Scan(&I)
			fmt.Scan(&N)
			fmt.Scan(&Q)
			ans = 0
			if I == Q {
				ans = N/2
			} else {
				ans = N - N/2
			}
			fmt.Println(ans)
		}
	}
}