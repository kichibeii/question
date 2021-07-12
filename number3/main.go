package main

import (
	"fmt"
	"math"
)

func main() {
	number := 262144
	counterValue := 0
	for i := 1; i <= number; i++ {
		counterFactor := 0
		numberCheck := i

		for j := 1; j <= int(math.Sqrt(float64(numberCheck))); j++ {
			if numberCheck%j == 0 {
				counterFactor++
			}
		}

		if int(math.Floor(math.Sqrt(float64(i)))) == int(math.Ceil(math.Sqrt(float64(i)))) {
			if counterFactor == 3 {
				counterFactor--
			}
		}

		if counterFactor == 3 {
			counterValue++
		}

	}

	fmt.Println(counterValue)
}
