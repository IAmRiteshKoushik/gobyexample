package main

import(
    "cmp"
    "fmt"
    "slices"
)

func sortByFunc() {
    // Attempt 1
    fruits := []string{ "peach", "banana", "kiwi" }
    lenCmp := func(a, b string) int {
        return cmp.Compare(len(a), len(b))
    }
    slices.SortFunc(fruits, lenCmp)
    fmt.Println(fruits)

    // Attempt 2
    type Person struct{
        name    string
        age     int
    }
    people := []Person{
        Person{name: "Jax", age: 37},
        Person{name: "TJ", age: 25},
        Person{name: "Alex", age: 72},
    }
    slices.SortFunc(people, func(a, b Person) int {
        return cmp.Compare(a.age, b.age)
    })
    fmt.Println(people)
}
