package main

import (
    "errors"
    "fmt"
)

func main() {
    result, err := divise(20.0, 0)

    if err != nil {
        panic(err)
    }

    fmt.Println(result)
}

func divise(a, b float32) (float32, error) {
    if b == 0 {
        return 0, errors.New("division par zéro impossible")
    }

    return a / b, nil
}
