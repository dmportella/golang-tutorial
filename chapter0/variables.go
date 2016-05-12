package chapter0

import "fmt"

func Variables() {
    var a string = "initial" // full hand
    var b, c int = 1, 2 // multi init
    var d = true // type inference
    var e int // init zero valued
    f := "short" // short hand (mostly used)

    const meh string = "some value"

    const num = 500000000

    fmt.Println(int64(num))

    fmt.Println(a, b, c, d, e, f, meh, "lets avoid compiler errors")
}