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

type UsersResult struct {
	TotalCount int `json:"total_count"`
	Items      []*User
}

type User struct {
	Login   string
	Id      int
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Title       string
	HTMLURL     string `json:"html_url"`
	Description string
	State       string
}

var bugsList = template.Must(template.New("bugsList").Parse(`
<h1>{{.TotalCount}} bug reports</h1>
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

var usersList = template.Must(template.New("usersList").Parse(`
<h1>{{.TotalCount}} users</h1>
<table>
<tr style='textalign:left'>
<th>Login</th>
<th>Id</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
<td>{{.Id}}</td>
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

/* Iterate over all the issues, fill in the map of users, */
/* and copy them into UsersResult structure */
func getUniqueUsers() *UsersResult {
	usersMap := make(map[int]*User)
	for _, issue := range total_issues.Items {
		curUserId := issue.User.Id

		if _, ok := usersMap[curUserId]; !ok {
			usersMap[curUserId] = issue.User
		}
	}

	var usersResult UsersResult
	usersResult.TotalCount = len(usersMap)

	users := make([]*User, 0, len(usersMap))
	for _, user := range usersMap {
		users = append(users, user)
	}
	usersResult.Items = users

	return &usersResult
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := getUniqueUsers()
	if err := usersList.Execute(w, users); err != nil {
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
	http.HandleFunc("/users", UsersHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
