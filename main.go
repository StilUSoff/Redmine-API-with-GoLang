package main

import (
	"flag"
	"fmt"
	"os"

	redmine "github.com/nixys/nxs-go-redmine/v2"
)

func main() {

	var r redmine.Context

	rdmnHost := flag.String("Host", "", "")
	rdmnAPIKey := flag.String("API_Key", "", "")
	flag.Parse()

	// Get variables from environment for connect to Redmine server
	if *rdmnHost == "" || *rdmnAPIKey == "" {
		fmt.Println("Init error: make sure you write variables `Host` and `API_Key` in file start")
		os.Exit(1)
	}

	// Init Redmine ctx
	r.SetEndpoint(*rdmnHost)
	r.SetAPIKey(*rdmnAPIKey)

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
}
