package main

import (
    "fmt"

    "./sequence"
)

func main() {
    square := func(x int) int { return x * x }
    s1 := sequence.StreamMap(square, sequence.MakeInt(1, 10))

    /*
     * rangeはループを回す要素数分繰り返すために使う
     */
    for x := range s1 {
        fmt.Print(x, " ")
    }

    fmt.Println()

    isOdd := func (x int) bool { return x % 2 != 0 }
    s2 := sequence.StreamFilter(isOdd, sequence.MakeInt(1, 20))
    for x := range s2 {
        fmt.Print(x, " ")
    }

    fmt.Println()
}
