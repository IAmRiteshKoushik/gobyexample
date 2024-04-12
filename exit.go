package main

import(
    "fmt"
    "os"
)

func exitProg() {
    defer fmt.Println("!")

    // C programs use an exit code which is basically the return value 
    // from the main function (for int main types)

    // If a similar behavior has to be created in Go, then os.Exit() must
    // be used. This will also make sure that the defer function is never
    // called. 
    os.Exit(3)

    // If "go run" is used then the exit status is printed on console
    // If "go build" is used then "echo $?" will show the exit code
}
