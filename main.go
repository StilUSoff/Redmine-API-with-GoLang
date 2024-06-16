package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	redmine "github.com/nixys/nxs-go-redmine/v2"
)

var r redmine.Context

func initRedmine() {
	rdmnHost := os.Getenv("HOST")
	rdmnAPIKey := os.Getenv("API_Key")

	if rdmnHost == "" || rdmnAPIKey == "" {
		fmt.Println("Init error: make sure you write variables `HOST` and `API_Key` in environment")
		os.Exit(1)
	}

	r.SetEndpoint(rdmnHost)
	r.SetAPIKey(rdmnAPIKey)

	fmt.Println("Init: success")
}

func getProjects(w http.ResponseWriter, req *http.Request) {
	p, _, err := r.ProjectAllGet([]string{"trackers", "issue_categories", "enabled_modules"})
	if err != nil {
		http.Error(w, fmt.Sprintf("Projects get error: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(p.Projects)
}

func createIssue(w http.ResponseWriter, req *http.Request) {
	projectID, err := strconv.Atoi(req.URL.Query().Get("ProjectID"))
	if err != nil {
		http.Error(w, "Invalid ProjectID", http.StatusBadRequest)
		return
	}

	trackerID, err := strconv.Atoi(req.URL.Query().Get("TrackerID"))
	if err != nil {
		http.Error(w, "Invalid TrackerID", http.StatusBadRequest)
		return
	}

	subject := req.URL.Query().Get("Subject")
	description := req.URL.Query().Get("Description")

	issue := redmine.IssueCreateObject{
		ProjectID:   projectID,
		Subject:     subject,
		Description: description,
		TrackerID:   trackerID,
	}

	createdIssue, _, err := r.IssueCreate(issue)
	if err != nil {
		http.Error(w, fmt.Sprintf("Issue creation error: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdIssue)
}

func main() {
	initRedmine()

	http.HandleFunc("/view_projects", getProjects)
	http.HandleFunc("/create_issue", createIssue)

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
