package main

import(
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func sFormat() {
    p := point{1, 2}    
    fmt.Printf("struct1: %v\n", p)
    fmt.Printf("struct2: %+v\n", p)
    fmt.Printf("struct3: %#v\n", p)

    fmt.Printf("type: %T\n", p)
    fmt.Printf("bool: %t\n", true)
    fmt.Printf("int: %d\n", 123)
    fmt.Printf("bin: %b\n", 14)
    fmt.Printf("char: %c\n", 33)
    fmt.Printf("hex: %x\n", 456)

    fmt.Printf("float1: %f\n", 78.9) // 6 digit precision
    fmt.Printf("float2: %e\n", 123400000.0) // 
    fmt.Printf("float3: %E\n", 123400000.0) // 

    fmt.Printf("str1: %s\n", "\"string\"") // String insert
    fmt.Printf("str2: %q\n", "\"string\"") // No format insert
    fmt.Printf("str3: %x\n", "hex this") // Hex insert

    fmt.Printf("pointer: %p\n", &p)

    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)

    // You can format anr print to io.Writers other 
    // than os.Stdout using Fprintf
    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
