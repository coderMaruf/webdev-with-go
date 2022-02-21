package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age int
	Hobbies [] string
}
// const data = `
// 	Hey, I am {{$name := .Name}}. I am {{.Age}} years old.{{range .Hobbies}} 
// 	{{$name}} like {{.}}{{end}}
// `
const data = `{{$name := .Name}}{{range .Hobbies}} 
{{$name}} like {{.}}{{end}}
`

func main() {
	maruf := Person{
		Name: "Maruf Sarker", 
		Age: 21,
		Hobbies: [] string{"Programming", "Web Design", "problem solving", "Tour"},
	}

	template := template.New("Maruf")

	template, err := template.Parse(data)
	if err != nil {
		log.Fatal("Parse error: ", err)
		return
	}
	
	template.Execute(os.Stdout, maruf)
}
