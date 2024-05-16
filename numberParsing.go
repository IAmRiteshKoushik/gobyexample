package main

import(
    "fmt"
    "strconv"
)

func numParse(){
    // Parse as a float and store in 64 bits
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    u, _ 
}
