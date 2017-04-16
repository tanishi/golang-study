package main

import "fmt"

func fibonacci() func() int {
    var fib [10]int

    fib[0] = 0
    fib[1] = 1

    for i := 2; i < 10; i++ {
        fib[i] = fib[i - 2] + fib[i - 1]
    }


    i := -1

    return func() int {
        i++

        return fib[i]
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}

