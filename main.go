package main

import (
	"fmt"
	"os"

	redmine "github.com/nixys/nxs-go-redmine/v2"
)

func main() {

	var r redmine.Context

	rdmnHost := os.Getenv("HOST")
	rdmnAPIKey := os.Getenv("API_Key")

	// Get variables from environment for connect to Redmine server
	if rdmnHost == "" || rdmnAPIKey == "" {
		fmt.Println("Init error: make sure you write variables `HOST` and `API_Key` in file start")
		os.Exit(1)
	}

	// Init Redmine ctx
	r.SetEndpoint(rdmnHost)
	r.SetAPIKey(rdmnAPIKey)

	fmt.Println("Init: success")

	// Get all projects
	p, _, err := r.ProjectAllGet([]string{"trackers", "issue_categories", "enabled_modules"})
	if err != nil {
		fmt.Println("Projects get error:", err)
		os.Exit(1)
	}

	fmt.Println("Projects:")
	for _, e := range p.Projects {
		fmt.Println("-", e.Name)
	}

	// Create a new issue
	issue := redmine.IssueCreateObject{
		ProjectID:   1, // Replace with an actual project ID
		Subject:     "New Issue from Go",
		Description: "This is a test issue created from Go",
		TrackerID:   1, // Replace with an actual tracker ID
	}

	createdIssue, _, err := r.IssueCreate(issue)
	if err != nil {
		fmt.Println("Issue creation error:", err)
		os.Exit(1)
	}

	fmt.Println("Issue created successfully with ID:", createdIssue.ID)
}
