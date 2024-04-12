package main

import "fmt"

func intSeq() func() int {
    i := 0
    // Anonymous function
    return func() int {
        i++ 
        return i
    }
}

func closure(){
    // Initialize a closure()
    // This returns a function so nextInt is basically a reference to a function
    // which has a return value of "i". Calling nextInt() again and again means
    // calling that function again and again and printing out the result

    nextInt := intSeq() // nextInt = () = {} (like javascript)

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // Initialize a closure()
    newInts := intSeq()
    fmt.Println(newInts())
}
