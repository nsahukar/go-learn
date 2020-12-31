package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("template-white.gohtml"))
}

func main() {
	home := Page{
		Title:   "Escaped",
		Heading: "Danger is escaped with html/template",
		Input:   `<script>alert("Yow!");</script>`,
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "template-white.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}
