package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    i := 0

    for {
        if math.Abs(z - (math.Pow(z, 2) - x) / (2 * z)) <= 0.001 {
            break;
        }
        z = z - (math.Pow(z, 2) - x) / (2 * z)
        i++
        fmt.Println(math.Abs(z - (math.Pow(z, 2) - x) / (2 * z)))
    }

	fmt.Println(i)

    return z
}

func main() {
	fmt.Println(Sqrt(2))
}

