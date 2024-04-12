package main

import (
    // Go implements several hash functions inside its crypto library
	"crypto/sha256"
	"fmt"
)

func shasha() {
    // Simple string
    s := "sha256 this string"
    
    // new hash function which returns an interface
    h := sha256.New()
    // Write expects bytes so if you have string, then convert it
    h.Write([]byte(s))
    // Can be used to append things to an existing hash slice
    bs := h.Sum(nil)

    // Original string
    fmt.Println(s)
    // Hashed string
    fmt.Printf("%x\n", bs)
}
