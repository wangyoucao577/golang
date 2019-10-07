package main

import (
	"github.com/wangyoucao577/golang_test/my_gopl_test/ch4_github"
	"log"
	"os"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)
// if err != nil {
// 	log.Fatal(err)
// }
var report = template.Must(template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
