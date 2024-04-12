package main

import(
    "fmt"
    "slices"
)

func slicey(){
    // Empty slice with length 0
    var s []string 
    fmt.Println("unint:", s, s == nil, len(s) == 0)

    // Empty slice with length 3 and capacity 3 even though cap is not specified
    s = make([]string, 3)
    fmt.Println("emp:", s, "len: ", len(s), "cap: ", cap(s))

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"

    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // Slice 1
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("sl1", c)

    // Slice 2
    l := s[:5]
    fmt.Println("sl2", l)

    // Slice 3
    l = s[2:]
    fmt.Println("sl3", l)

    // Slice 4
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // Slice 5
    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2){
        fmt.Println("t == t2")
    }

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
