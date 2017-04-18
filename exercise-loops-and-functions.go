package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	i := 0

	for {
		if math.Abs(z-(z-(math.Pow(z, 2)-x)/(2*z))) <= 1e-10 {
			break
		}
		z = z - (math.Pow(z, 2)-x)/(2*z)
		i++
	}

	fmt.Println(i)

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
