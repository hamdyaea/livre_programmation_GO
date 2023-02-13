package main

import ("fmt")

func PeriRect(long,larg int) int{
    return 2 * (long + larg)
}

func main() {
    rectangle := PeriRect(40,70)
    fmt.Printf("Le périmètre du rectangle est : %v. \n", rectangle)
}
