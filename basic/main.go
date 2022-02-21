package main

import (
	"text/template"
	"log"
	"os"
)

type Person struct {
	Name string
}

func main() {
	name := Person{"Maruf Sarker"}
	
	template := template.New("test")

	template, err := template.Parse("Hello, {{.Name}} \n")

		if err != nil {
			log.Fatal("Parse error: ", err)
			return
		}
	template.Execute(os.Stdout, name, )
}
