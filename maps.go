package main

import(
    "fmt"
    "maps"
)

func mapEx(){
    // Initializing a map
    m := make(map[string]int)

    // Adding key-value pairs to a map
    m["k1"] = 7
    m["k2"] = 8
    // Printing out the entire map
    fmt.Println("map:", m)

    // Getting the value from a key
    v1 := m["k1"]
    fmt.Println("v1:", v1)

    // Getting another value from the key (which doesn't exist)
    v3 := m["k3"]
    fmt.Println("v3:", v3)

    // Printing the entire length of the map
    fmt.Println("len:", len(m))

    // Deleting the key-value pair from the map
    delete(m, "k2")
    fmt.Println("map:", m)

    // Clear the entire map but not deleting it
    clear(m)
    fmt.Println("map:", m)

    // Checking the presence of an element
    // Syntactic sugar by Go

    // value, presence := map["key-here"]
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // Using map functions
    n := map[string]int {"foo": 1, "bar": 2}
    fmt.Println("map:", n)
    n2 := map[string]int {"foo": 1, "bar": 2}
    if maps.Equal(n, n2){
        fmt.Println("n == n2")
    }
}
