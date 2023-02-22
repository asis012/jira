package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/andygrunwald/go-jira"
)

const (
	JIRAURL = "https://ashishgyawalijira.atlassian.net/"
)

func main() {

	// Replace with your Jira API token
	token := ""
	// Replace with your Jira username
	username := "ashishgyawali@lftechnology.com"

	// Set up an HTTP client
	client := &http.Client{}

	// Create a new request to the Jira API
	req, err := http.NewRequest("GET", JIRAURL+"/rest/api/2/myself", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the authorization header with the token and username
	req.SetBasicAuth(username, token)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the response JSON
	var user map[string]interface{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func integrateGoJira() {
	tp := jira.BasicAuthTransport{
		Username: "username",
		Password: "token",
	}

	jiraClient, err := jira.NewClient(tp.Client(), JIRAURL)
	if err != nil {
		panic(err)
	}

	issue, _, _ := jiraClient.Issue.Get("JIR-1", nil)
	currentStatus := issue.Fields.Status.Name
	fmt.Printf("Current status: %s\n", currentStatus)

	var transitionID string
	possibleTransitions, _, _ := jiraClient.Issue.GetTransitions("JIR-1")
	for _, v := range possibleTransitions {
		if v.Name == "In Progress" {
			transitionID = v.ID
			break
		}
	}

	jiraClient.Issue.DoTransition("JIR-1", transitionID)
	issue, _, _ = jiraClient.Issue.Get("JIR-1", nil)
	fmt.Printf("Status after transition: %+v\n", issue.Fields.Status.Name)
}
