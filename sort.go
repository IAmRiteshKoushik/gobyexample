package main

import(
    "fmt"
    "slices"
)

func sortIt() {
    // Soring functions are generic
    strs := []string{"c", "a", "b", "ab"}
    slices.Sort(strs)
    fmt.Println("Strings:", strs)

    // Ascending order sorting takes place
    ints := []int{7, 2, 4, 22}
    slices.Sort(ints)
    fmt.Println("Ints:   ", ints)

    s := slices.IsSorted(ints)
    fmt.Println("Sorted: ", s)
}
