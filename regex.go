package main

import(
    "bytes"
    "fmt"
    "regexp"
)

func RegexCheck() {
    // -- STRINGS
    // Testing whether a pattern matches a string
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    // Compiling an optimized regex struct - many methods are available
    // "Reusable" regex check
    r, _ := regexp.Compile("p([a-z]+)ch")

    // Normal regex check against a string
    fmt.Println(r.MatchString("peach"))
    // Returns the regex string which matches
    fmt.Println(r.FindString("peach punch"))
    // Finds the first match and returns the start and end indexes for the match
    fmt.Println("idx:", r.FindStringIndex("peach punch"))
    // Returns information on "p([a-z]+)ch", "([a+z]+)"
    fmt.Println(r.FindStringSubmatch("peach punch"))
    
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))
   
    fmt.Println(r.FindAllString("peach punch pinch", -1))
    // Setups up no limits to the number of matches if negative arg is present
    fmt.Println("all", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
    // Limits the number of matches to 2
    fmt.Println(r.FindAllString("peach punch pinch", 2))



    // -- BYTES INSTEAD OF STRINGS 
    // Instead of string, you can also supply byte arguments and drop the 
    // "string" from the method name to parse byte content
    fmt.Println(r.Match([]byte("peach")))

    // Replace subsets of strings with other values
    r = regexp.MustCompile("p([a-z]+)ch")
    // Prints out the regex string
    fmt.Println("regexp:", r)
    // Replaces all occurances of a particular regex satisfied with alt text
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    // Transform matches text using a function as an argument
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}
