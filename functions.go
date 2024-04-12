package main

import "fmt"

func plus(a, b int) int {
    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func multiReturn() (int, int) {
    return 3, 7
}

func callFunc() {
    
    res := plus(1, 2)
    fmt.Println("1 + 2 = ", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1 + 2 + 3 = ", res)

    _, c := multiReturn()
    fmt.Println(c)
}
