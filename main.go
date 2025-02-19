package main

import (
	"fmt"
	"simple_git_activity/internal/parser"
)

func main() {
	var URL string = "https://api.github.com/users/"
	var name string
	fmt.Print("User name : ")
	fmt.Scanln(&name)
	if len(name) == 0 {
		fmt.Println("User name is empty")
	}
	var fullURL string = URL + name + "/events"

	events, err := parser.GetGitActivity(fullURL)
	if err != nil {
		fmt.Println(err)
	}
	parser.PrintGitEvents(events)
}
