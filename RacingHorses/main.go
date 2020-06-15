package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var count int
		fmt.Scan(&count)
		HorseSkills := make([]int, count)

		for i := 0; i< count; i++ {
			fmt.Scan(&HorseSkills[i])

			j := i - 1
			key := HorseSkills[i]
			for ; j>= 0 && HorseSkills[j] > key; {
				HorseSkills[j+1] = HorseSkills[j]
				j--
			}
			HorseSkills[j + 1] = key
		}
		var diff = HorseSkills[1] - HorseSkills[0]
		for i := 2; i < count; i++ {
			if HorseSkills[i] - HorseSkills[i - 1] < diff {
				diff = HorseSkills[i] - HorseSkills[i -1]
			}
		}
		fmt.Println(diff)
	}
}
