package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	i := 0

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	for {
		if math.Abs(z-(z-(math.Pow(z, 2)-x)/(2*z))) <= 1e-10 {
			break
		}
		z = z - (math.Pow(z, 2)-x)/(2*z)
		i++
	}

	fmt.Println(i)

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
