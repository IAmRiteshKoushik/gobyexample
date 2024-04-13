package main

import(
    "fmt"
    "os"
)

// Main function
func BegToDefer() {
    f := createFile("/home/rk/defer.txt")
    defer closeFile(f) // Executed at the end of the main function
    // Basically the execution of the function has been delayed until later
    writeFile(f)
}

// Helper functions
func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing") 
    fmt.Fprintln(f, "data")
}

// Checking for errors in a deferred function
func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}
