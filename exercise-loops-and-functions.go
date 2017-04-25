package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	eps := 1e-10

	for {
		if math.Abs(z-(z-(math.Pow(z, 2)-x)/(2*z))) <= eps {
			break
		}
		z = z - (math.Pow(z, 2)-x)/(2*z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
