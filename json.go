package main 

import(
    "encoding/json"
    "fmt"
    // "os"
)

// Encoding a JSON
type response1 struct {
    Page int
    Fruits []string
}

// Decoding a JSON 
type response2 struct {
    Page    int         `json:"page"`
    Fruits  []string    `json:"fruits"`
}

// Only exported fields are encoded/decoded in JSON,
// and they must start with capital letters to be exported

func jsonEx() {
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))
    fmt.Println(bolB) // Converts into an array of equaivalent decimals

    intB, _ := json.Marshal(10)
    fmt.Println(string(intB))
    fmt.Println(intB) // Converts to array of decimal of the character 

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))
    fmt.Println(fltB)

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))
    fmt.Println(strB)
}
