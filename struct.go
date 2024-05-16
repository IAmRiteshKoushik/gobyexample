package main

import "fmt"

// Declaring a struct
type person struct {
    name string
    age int
}

// Acts like a constructor method
func newPerson(name string) *person {
    p := person{ name: name }
    p.age = 42
    return &p // Returning a pointer because it will survive the scope of func
}

func structs() {

    // Preliminary trials
    fmt.Println(person{ "Bob", 20 })
    fmt.Println(person{name: "Alice", age: 30})
    fmt.Println(person{name: "Fred"})
    fmt.Println(&person{name: "Ann", age: 40})

    // Using constructor-like method
    fmt.Println(newPerson("Jon"))

    // Trail 1
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // Trial 2
    sp := &s
    fmt.Println(sp.age)
    
    // Trial 3
    sp.age = 51
    fmt.Println(sp.age)

    // Struct declared and initalized in single line
    dog := struct {
        name    string
        isGood  bool
    }{ "Rex", true}
    fmt.Println(dog)
}
