package main

import "fmt"

// Simple function
func zeroval(ival int) {
    ival = 0
}

// Function with pointer argument
func zeroptr (iptr *int) {
    // Setting the value at the address to 0
    *iptr = 0
}

func pointers () {
    i := 1
    fmt.Println("Initial:", i)

    zeroval(i)
    // No change in original variable
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    // Change in original variable
    fmt.Println("zeroptr:", i)

    // Address of the pointers
    fmt.Println("pointer", &i)
}
