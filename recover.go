package main

import "fmt"

func mayPanic() {
    panic("A problem has occured")
}

func LetsRecover() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered, Error:\n", r)
        }
        // Expression in defer must be a function call
    }()
    mayPanic()
    fmt.Println("After mayPanic()")
}
