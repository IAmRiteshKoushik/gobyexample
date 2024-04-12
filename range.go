package main

import "fmt"

func rangeEx(){
    nums := []int{2, 3, 4}
    sum := 0
    // Normally iterating over an array
    for i := range nums {
        fmt.Println(i)
        // By default only index is accessed and not the element
        // Basically, you have not captured theh entire result returned
        // by the range function
    }

    // Iterating over an array with : _, element syntax
    // Printing out array element
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // Iterating over an array with : index, element syntax
    for i, num := range nums {
        if num == 3 {
            fmt.Println("Index:", i)
        }
    }

    kvs := map[string]string {"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    for k := range kvs {
        fmt.Println("key:", k)
    }
    for i, c := range "go" {
        fmt.Println(i, c)
        // Prints out the decimal equivalents from the ASCII table
    }
}
