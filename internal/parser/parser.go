package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
		Action string `json:"action,omitempty"`
		Issue  struct {
			Title  string `json:"title"`
			Labels []struct {
				Id          int         `json:"id"`
				NodeId      string      `json:"node_id"`
				Url         string      `json:"url"`
				Name        string      `json:"name"`
				Color       string      `json:"color"`
				Default     bool        `json:"default"`
				Description interface{} `json:"description"`
			} `json:"labels"`
			State             string        `json:"state"`
			Locked            bool          `json:"locked"`
			Assignee          interface{}   `json:"assignee"`
			Assignees         []interface{} `json:"assignees"`
			Milestone         interface{}   `json:"milestone"`
			Comments          int           `json:"comments"`
			CreatedAt         time.Time     `json:"created_at"`
			UpdatedAt         time.Time     `json:"updated_at"`
			ClosedAt          time.Time     `json:"closed_at"`
			AuthorAssociation string        `json:"author_association"`
			SubIssuesSummary  struct {
				Total            int `json:"total"`
				Completed        int `json:"completed"`
				PercentCompleted int `json:"percent_completed"`
			} `json:"sub_issues_summary"`
			ActiveLockReason interface{} `json:"active_lock_reason"`
			Body             string      `json:"body"`
			Reactions        struct {
				Url        string `json:"url"`
				TotalCount int    `json:"total_count"`
				Field3     int    `json:"+1"`
				Field4     int    `json:"-1"`
				Laugh      int    `json:"laugh"`
				Hooray     int    `json:"hooray"`
				Confused   int    `json:"confused"`
				Heart      int    `json:"heart"`
				Rocket     int    `json:"rocket"`
				Eyes       int    `json:"eyes"`
			} `json:"reactions"`
			TimelineUrl           string      `json:"timeline_url"`
			PerformedViaGithubApp interface{} `json:"performed_via_github_app"`
			StateReason           *string     `json:"state_reason"`
			Draft                 bool        `json:"draft,omitempty"`
			PullRequest           struct {
				Url      string     `json:"url"`
				HtmlUrl  string     `json:"html_url"`
				DiffUrl  string     `json:"diff_url"`
				PatchUrl string     `json:"patch_url"`
				MergedAt *time.Time `json:"merged_at"`
			} `json:"pull_request,omitempty"`
		} `json:"issue,omitempty"`
		PullRequest struct {
			State string `json:"state"`
			Title string `json:"title"`
		} `json:"pull_request,omitempty"`
	} `json:"payload"`
}

// GetGitActivity Parse git api
func GetGitActivity(URL *string) error {
	// getting url from usr
	resp, err := http.Get(*URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// reading body response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// creating struct
	var events []GitApi
	// unmarshal a body
	err = json.Unmarshal(body, &events)
	if err != nil {
		log.Fatal(err)
	}

	// loop range for printing events
	for _, event := range events {
		commitCount := len(event.Payload.Commits)

		switch event.Type {
		case "PushEvent":
			fmt.Printf("Pushed %v commits to %s\n", commitCount, event.Repo.Name)
			for _, commit := range event.Payload.Commits {
				fmt.Printf("\tCommits: %s\n", commit.Message)
			}
		case "IssueCommentEvent":
			fmt.Printf("Opened a new issue in %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("Started %s\n", event.Repo.Name)
		case "CreateEvent":
			fmt.Printf("Created new repo %s\n", event.Repo.Name)
		}
	}

	return nil
}
