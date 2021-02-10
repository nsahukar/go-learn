package main

import (
	"json-decode/github"
	"log"
	"os"
	"text/template"
	"time"
)

const tmpl = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------------------
Number: {{.Number}}
User: 	{{.User.Login}}
Title: 	{{.Title | printf "%.64s"}}
Age: 	{{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(tmpl))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Printf("%d issues:\n", result.TotalCount)
	//	for _, item := range result.Items {
	//		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	//	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
