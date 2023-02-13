package main


import (
    "fmt"
    "monProgramme/mesTypes"
)



func main() {

    var testTypes mesTypes.Types

    testTypes.TypeI = 24
    testTypes.TypeStr = "Bonjour"
    testTypes.TypeB = true

    fmt.Println(testTypes.TypeI)

    fmt.Println(testTypes)

}
