package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func newton(x float64) func() float64 {
	z := 1.0

	return func() float64 {
		r := z

		z = z - (math.Pow(z, 2)-x)/(2*z)

		return r
	}
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	n := newton(x)

	current := n()

	for {
		next := n()

		if math.Abs(current-next) <= 1e-10 {
			return current, nil
		}

		current = next
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
