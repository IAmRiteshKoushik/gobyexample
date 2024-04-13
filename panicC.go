package main

import "os"

func panicC() {

    // Panicing means not handling the error gracefully
    // Common use of panic is toabort if a function returns an error
    // value that we do not know how to (or want to) handle.
    panic("a problem")
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}

// While other langauges use exceptions for handling many errors, in 
// Go it is idiomatic to use error-indicating return values wherever possible
