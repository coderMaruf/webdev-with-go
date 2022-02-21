package main

import (
	"text/template"
	"log"
	"os"
)

func main() {
	template, err := template.ParseFiles("template.gohtml")

	if err != nil {
		log.Fatal(err)
		return
	}
	template.Execute(os.Stdout, nil)
}
