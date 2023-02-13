package main

import (
    "fmt"
    "math/rand"
    "time"
)

const (
    letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    specialBytes  = "!@#$;?"
    digitBytes    = "0123456789"
    passwordLen   = 16
    letterIdxMask = 1<<6 - 1
    specialIdxMask = 1<<7 - 1
    digitIdxMask = 1<<5 - 1
)

func main() {
    rand.Seed(time.Now().UnixNano())
    password := generatePassword(passwordLen)
    fmt.Println("Votre mot de passe est :", password)
}

func generatePassword(n int) string {
    password := make([]byte, n)
    for i, cache, remain := n-1, rand.Int63(), passwordLen; i >= 0; {
        if remain == 0 {
            cache, remain = rand.Int63(), passwordLen
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            password[i] = letterBytes[idx]
            i--
        }
        cache >>= 6
        remain--

        if remain == 0 {
            cache, remain = rand.Int63(), passwordLen
        }
        if idx := int(cache & specialIdxMask); idx < len(specialBytes) {
            password[i] = specialBytes[idx]
            i--
        }
        cache >>= 7
        remain--

        if remain == 0 {
            cache, remain = rand.Int63(), passwordLen
        }
        if idx := int(cache & digitIdxMask); idx < len(digitBytes) {
            password[i] = digitBytes[idx]
            i--
        }
        cache >>= 5
        remain--
    }
    return string(password)
}

