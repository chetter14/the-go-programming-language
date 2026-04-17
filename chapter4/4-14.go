package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
)

/* Create a web server that queries GitHub once and then allows navigation of the
list of bug reports, milestones, and users. */

const IssuesURL = "https://api.github.com/search/issues?q=repo:golang/go"

type IssuesResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
	Labels    []*Label
	Milestone *Milestone
}

type Label struct {
	Name string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Title       string
	HTMLURL     string `json:"html_url"`
	Description string
	State       string
}

var bugsList = template.Must(template.New("bugsList").Parse(`
<h1>{{.TotalCount}} issues (bug reports)</h1>
<table>
<tr style='textalign:left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

/* Find only bugs in all the issues */
func findBugs() *IssuesResult {
	var bugs IssuesResult
	for _, issue := range total_issues.Items {
		/* Search through this issue labels */
		for _, label := range issue.Labels {
			if label.Name == "BugReport" {
				bugs.TotalCount++
				bugs.Items = append(bugs.Items, issue)
			}
		}
	}
	return &bugs
}

func BugReportsHandler(w http.ResponseWriter, r *http.Request) {
	bugs := findBugs()
	if err := bugsList.Execute(w, bugs); err != nil {
		log.Fatal(err)
	}
}

var total_issues IssuesResult

func main() {
	resp, err := http.Get(IssuesURL)
	if err != nil {
		log.Fatalf("failed to get the issues: %s", resp.Status)
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("failed to query all the issues: %s", resp.Status)
		return
	}

	if err = json.NewDecoder(resp.Body).Decode(&total_issues); err != nil {
		resp.Body.Close()
		log.Fatalf("failed to unmarshall the issues: %s", resp.Status)
		return
	}

	resp.Body.Close()

	http.HandleFunc("/bugs", BugReportsHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
