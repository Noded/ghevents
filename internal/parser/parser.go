package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	pushEvent         = "PushEvent"
	issueCommentEvent = "IssueCommentEvent"
	watchEvent        = "WatchEvent"
	createEvent       = "CreateEvent"
)

// GitApi Its json format of api.github.com
type GitApi struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits,omitempty"`
		// Другие поля можно добавить по необходимости
	} `json:"payload"`
}

// GetGitActivity Parse git api
func GetGitActivity(URL string) ([]GitApi, error) {
	// getting url from usr
	resp, err := http.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("fail to get response: %s", err)
	}
	defer resp.Body.Close()

	// reading body response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fail read response: %s", err)
	}

	// creating struct
	var events []GitApi
	// unmarshal a body
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal json: %s", err)
	}

	return events, nil
}

func PrintGitEvents(events []GitApi) {
	// loop range for printing events
	for _, event := range events {
		commitCount := len(event.Payload.Commits)

		switch event.Type {
		case pushEvent:
			fmt.Printf("Pushed %v commits to %s\n", commitCount, event.Repo.Name)
			for _, commit := range event.Payload.Commits {
				fmt.Printf("\tCommits: %s\n", commit.Message)
			}
		case issueCommentEvent:
			fmt.Printf("Opened a new issue in %s\n", event.Repo.Name)
		case watchEvent:
			fmt.Printf("Started %s\n", event.Repo.Name)
		case createEvent:
			fmt.Printf("Created new repo %s\n", event.Repo.Name)
		}
	}
}
