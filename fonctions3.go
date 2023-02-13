package main

import (
    "fmt"
    "errors"
)

func DitAurevoir(name string) (string, error) {
    if name == "" {
        return "", errors.New("\bVous devez spécifier un nom")
    }
    message := fmt.Sprintf("%v part! Bonne journée.", name)
    return message,nil
}

func main() {
    fmt.Println(DitAurevoir(""))
}
