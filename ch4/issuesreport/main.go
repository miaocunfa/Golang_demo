package main

import (
	"../github"
	"html/template"
	"log"
	"os"
	"time"
)

//const templ = `{{.TotalCount}} issues:
//{{range .Items}}--------------------------------
//Number: {{.Number}}
//User:   {{.User.Login}}
//Title:  {{.Title | printf "%.64s"}}
//Age:    {{.CreateAt | daysAgo }} days
//{{end}}
//`

const htmltempl = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`


func daysAgo(t time.Time) int  {
	return int(time.Since(t).Hours() / 24)
}

//var report = template.Must(template.New("issuelist").
//	Funcs(template.FuncMap{"daysAgo": daysAgo}).
//	Parse(templ))

var htmlreport = template.Must(template.New("htmlissue").Parse(htmltempl))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	//if err := report.Execute(os.Stdout, result); err != nil {
	//	log.Fatal(err)
	//}

	if err := htmlreport.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
