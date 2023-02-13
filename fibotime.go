package main

import (
  "fmt"
  "time"
)

func fib(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else {
        return fib(n-1) + fib(n-2)
    }
}

func main() {
    start := time.Now()

    for i := 0; i <= 45; i++ {
        fmt.Printf("Calculating Fibonacci number %d...\n", i)
        result := fib(i)
        elapsed := time.Since(start)
        fmt.Printf("Fibonacci number %d is %d. Calculated in %s.\n", i, result, elapsed)
    }
}
