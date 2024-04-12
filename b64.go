package main

import(
    b64 "encoding/base64"
    "fmt"
)

func baseAndURLEnc() {

    data := "abc123!?$*&()'-=@~"

    // Standard encoding
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc) // Encoded string
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(sDec) // Actual format
    fmt.Println(string(sDec)) // String format

    // Standard decoding
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc) // Encoded string
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(uDec) // Actual format
    fmt.Println(string(uDec)) // String format
} 
