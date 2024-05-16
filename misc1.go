package main

import "fmt"

func misc1() {
    const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

    // Printing out the entire sample
    fmt.Println("Println: ")
    fmt.Println(sample)

    // Trial 0
    fmt.Println("ByteLoop: ")
    for i := 0; i < len(sample); i++ {
        fmt.Printf("%x\n", sample[i])
    }
    fmt.Printf("\n")

    // Trial 1
    fmt.Println("Printf with %x:")
    fmt.Printf("%x\n", sample)

    // Trial 2
    fmt.Println("Printf with % x:")
    fmt.Printf("% x\n", sample)

    // Trial 3
    fmt.Println("Printf with %q:")
    fmt.Printf("%q\n", sample)

    // Trial 4
    fmt.Println("Printf with %+q:")
    fmt.Printf("%+q\n", sample)
}

func misc2() {

}
