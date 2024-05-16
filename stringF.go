package main

import(
    str "strings"
    // There is a build in s variable inside strings
    "fmt"
)

// Aliasing 
var p = fmt.Println

func stringF() {
    p("Contains:  ", str.Contains("test", "es"))
    p("Count:     ", str.Count("test", "t"))

    p("HasPrefix: ", str.HasPrefix("test", "te"))
    p("HasSuffix: ", str.HasSuffix("test", "st"))

    p("Index:     ", str.Index("test", "e"))
    p("Repeat:    ", str.Repeat("a", 5))

    p("Replace    ", str.Replace("foo", "o", "0", -1)) // Replaces all 
    p("Replace    ", str.Replace("foo", "o", "0", 1)) // Replaces first

    p("Split      ", str.Split("a-b-c-d-e", "-"))
    p("Join:      ", str.Join([]string{"a", "b"}, "-"))

    p("ToLower:   ", str.ToLower("TEST"))
    p("ToUpper:   ", str.ToUpper("test"))
}
