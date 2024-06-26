package main

import(
    "fmt"
    "unicode/utf8"
)

// Decoding each rune character
func examineRune(r rune){
    if r == 't' {
        fmt.Println("Found tee")
    } else if r == 'ส' {
        fmt.Println("Found so sua")
    }
}

// Driver code for runes and strings
func stringNRunes() {
    const s = "สวัสดี"
    fmt.Println("Len:", len(s))

    // 
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x", s[i])
    }
    fmt.Println()
    fmt.Println("Rune count:", utf8.RuneCountInString(s))

    // 
    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Println("%#U starts at %d\n", runeValue, i)
        w = width

        examineRune(runeValue)
    }
}
