package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"../github"
)

func daysAgo(t time.Time) int  {
	return int(time.Since(t).Hours() / 24)
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreateAt | daysAgo }} days
{{end}}
`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

/*
D:\git-work\Golang_demo\ch4\issuesreport>go build

D:\git-work\Golang_demo\ch4\issuesreport>issuesreport.exe repo:golang/go is:open json decoder
57 issues:
--------------------------------
Number: 43716
User:   ggaaooppeenngg
Title:  encoding/json: fix byte counter increments when using decoder.To
Age:    106751 days
--------------------------------
Number: 33416
User:   bserdar
Title:  encoding/json: This CL adds Decoder.InternKeys
Age:    106751 days
--------------------------------
...
 */