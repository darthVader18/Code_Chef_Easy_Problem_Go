package main
import (
	"fmt"
	"math"
)

func main() {
	var T, n, m int// n is the no. of notes in wallet and m is the amount of money mugger asked for
	fmt.Scan(&T)
	for ; T > 0; T-- {
		fmt.Scan(&n, &m)
		a := make([]int, n)
		for index := 0; index < n; index++ {
			fmt.Scan(&a[index])
		}
		z := int(math.Pow(float64(2), float64(n)))
		var flag bool = false
		for i := 0; i < z; i++ {
			j := i
			value := 0
			for p := 0; p < n; p++ {
				if j % 2 == 1 {
					value += a[p]
				}
				j >>= 1
			}
			if value == m {
				flag = true
				break
			}
		}
		if flag == true {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
