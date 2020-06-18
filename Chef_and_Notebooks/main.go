
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	inputs := strings.Fields(text)
	T,_ := strconv.ParseInt(inputs[0],10,64)

	for i := int64(0); i < T; i++ {
		text, _ = reader.ReadString('\n')
		inputs = strings.Fields(text)
		X,_ := strconv.ParseInt(inputs[0],10,64)//X is the no. of pages
		Y,_ := strconv.ParseInt(inputs[1],10,64)//only Y pages are left in the notebook. He can use all Y pages left in the notebook
		K,_ := strconv.ParseInt(inputs[2],10,64)//Chef has only K rubles(money) left for now
		N,_ := strconv.ParseInt(inputs[3],10,64)//Shopkeeper showed him N notebooks

		var result bool
		for j := int64(0); j < N; j++ {
			text, _ = reader.ReadString('\n')
			inputs = strings.Fields(text)
			P,_ := strconv.ParseInt(inputs[0],10,64)//No. of pages in the notebook
			C,_ := strconv.ParseInt(inputs[1],10,64)//Price

			if (P+Y) >= X && C <= K {
				result = true
			}
		}

		if result {
			fmt.Println("LuckyChef")
		} else {
			fmt.Println("UnluckyChef")
		}
	}
}


