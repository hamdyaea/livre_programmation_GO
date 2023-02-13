package main

import "fmt"

func main(){
    var (
        b bool
        s string
        i int
        u uint
        u8 uint8
        i8 int8
        i16 int16
        u16 uint16
        f float32

    )

    b = true
    s = "hello"
    i = -30
    u = 50
    u8 = 220
    i8 = -120
    i16 = 32000
    u16 = 50000
    f = 5.456

    fmt.Println(b, s, i, u, u8, i8, i16, u16, f)
}
