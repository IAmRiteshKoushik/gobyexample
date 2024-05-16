package main

// There is a package "html/" which can be used for generating HTML text
// like error messages : BadGateway, InternalServerError etc.

import(
    "os"
    "text/template"
)

func templateEx() {

    // Trial 1 - This is one way to define a template
    t1 := template.New("t1")
    t1, err := t1.Parse("Value is {{.}}\n")
    if err != nil {
        panic(err)
    }
    // Templates are a mix of static text and actions enclosed in {{...}}
    // This is another way of defining a template (we are using the same 
    // variable )
    t1 = template.Must(t1.Parse("Value: {{.}}\n"))
    t1.Execute(os.Stdout, "some text")
    t1.Execute(os.Stdout, 5)
    t1.Execute(os.Stdout, []string{
        "Go",
        "Rust",
        "C++",
        "C#",
    })

    // Anonymous function stored inside a variable placeholder
    Create := func(name, t string) *template.Template {
        return template.Must(template.New(name).Parse(t))
    }

    // Trial 2
    t2 := Create("t2", "Name: {{.Name}}\n")
    t2.Execute(os.Stdout, struct {
        Name string
    }{"Jane Doe"})
    t2.Execute(os.Stdout, map[string]string{"Name" : "Micky Mouse"})
    
    // Trial 3
    t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
    t3.Execute(os.Stdout, "not empty")
    t3.Execute(os.Stdout, "")

    // Trial 4
    t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
    t4.Execute(os.Stdout, []string{
        "Go",
        "Rust",
        "C++",
        "C#",
    })
}
