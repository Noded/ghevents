package main

import (
	"fmt"
	"simple_git_activity/internal/parser"
)

func main() {
	var URL string
	fmt.Print("URL: ")
	fmt.Scanln(&URL)

	parser.GetGitActivity(&URL)
}
