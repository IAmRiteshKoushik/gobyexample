package main

import(
    "fmt"
    "math"
)

const s string = "constant"

func constants(){
    fmt.Println(s)
    const n = 5000000

    // Arbitrary precision
    const d = 3e20 / n
    fmt.Println(d)

    // A number can be given a type by using it in 
    // a context that requires one, such as function
    // call
    fmt.Println(int64(d))
    fmt.Println(math.Sin(n))

}
