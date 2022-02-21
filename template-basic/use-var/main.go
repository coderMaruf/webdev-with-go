package main

import (
	"text/template"
	"log"
	"os"
)

func main() {
	var var1 = "Hello, world!"
	template, err := template.ParseFiles("template.gohtml")

	if err != nil {
		log.Fatal(err)
		return
	}
	template.ExecuteTemplate(os.Stdout, "template.gohtml", var1)
}
