package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n - 1)
}

func recursion(){
    // Simle recursion program
    fmt.Println(fact(7))

    // A variable which holds a reference to an anonymous function
    // This is similar to javascript's const <func-name> = () => {}
    var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n - 1) + fib(n - 2)
    }

    // Calling the function
    fmt.Println(fib(7))
}
