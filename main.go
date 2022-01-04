package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	var regex string = os.Getenv("PLUGIN_REGEX")
	var commitMessage = os.Getenv("DRONE_COMMIT_MESSAGE")

	fmt.Println("Regex: " + regex)
	fmt.Println("Commit Message: " + commitMessage)

	r, _ := regexp.Compile(regex)

	if r.MatchString(commitMessage) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
