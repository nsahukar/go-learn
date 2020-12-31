package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tmpl *template.Template

type user struct {
	Username string
	Name     string
	Email    string
	Phone    string
}

// create a template.FuncMap to register functions.
var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles(
		"template-white.gohtml",
		"template-variables.gohtml",
		"template-array.gohtml",
		"template-array-orange.gohtml",
		"template-map.gohtml",
		"template-map-orange.gohtml",
		"template-struct.gohtml",
		"template-struct-array.gohtml",
	))
}

func main() {
	// err := tmpl.ExecuteTemplate(os.Stdout, "template-white.gohtml", 42)
	// err := tmpl.ExecuteTemplate(os.Stdout, "template-variables.gohtml", "Release self-focus; Embrace other-focus")

	// colors := []string{"White", "Orange", "Blue", "Yellow", "Green", "Brown", "Black"}
	// err := tmpl.ExecuteTemplate(os.Stdout, "template-array-orange.gohtml", colors)

	// 	iataCodes := map[string]string{
	// 		"LAX": "Los Angeles International Airport",
	// 		"HND": "Tokyo Haneda Airport",
	// 		"ORD": "O'Hare International Airport",
	// 		"LHR": "London Heathrow Airport",
	// 		"CDG": "Charles de Gaulle Airport",
	// 		"AMS": "Amsterdam Airport Schiphol",
	// 		"HKG": "Hong Kong International Airport",
	// 		"ICN": "Seoul Incheon International Airport",
	// 		"FRA": "Frankfurt Airport",
	// 		"DEL": "Indira Gandhi International Airport",
	// 		"SIN": "Singapore Changi Airport",
	// 	}
	// 	err := tmpl.ExecuteTemplate(os.Stdout, "template-map-orange.gohtml", iataCodes)

	// nikhil := user{
	// 	Username: "nsahukar",
	// 	Name:     "Nikhil Sahukar",
	// 	Email:    "nsahukar@gmail.com",
	// 	Phone:    "+91 88883 88473",
	// }
	// err := tmpl.ExecuteTemplate(os.Stdout, "template-struct.gohtml", nikhil)

	nikhil := user{
		Username: "nsahukar",
		Name:     "Nikhil Sahukar",
		Email:    "nsahukar@gmail.com",
		Phone:    "+91 88883 88473",
	}
	dhruv := user{
		Username: "droov",
		Name:     "Dhruv Negi",
		Email:    "dhruvnegi@gmail.com",
		Phone:    "+91 81700 12275",
	}
	ankur := user{
		Username: "baani",
		Name:     "Ankur Bansal",
		Email:    "ankur.k.bansal@gmail.com",
		Phone:    "+91 99119 16101",
	}
	users := []user{nikhil, dhruv, ankur}
	err := tmpl.ExecuteTemplate(os.Stdout, "template-struct-array.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
