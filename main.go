package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Parse command-line arguments
	username := parseArgs()

	// Validate username from Github
	if err := validateUsername(username); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Fetching activity for GitHub user: %s\n", username)
}

func parseArgs() (username string) {
	help := flag.Bool("h", false, "Show help message")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	// first non flag argument, flag handles the program name
	username = args[0]

	// if with initialization
	if err := validateUsername(username); err != nil {
		invalidUser()
		os.Exit(1)
	}

	return username
}

func printUsage() {
	usage := `Usage: github-activity [options] <github-username>`
	fmt.Println(usage)
}

func invalidUser() {
	invalid := `Error: Invalid GitHub username.
`
	fmt.Println(invalid)
}

func validateUsername(username string) error {
	if username == "" {
		return fmt.Errorf("invalid username")
	}

	// Check if the username exists on GitHub
	resp, err := http.Get("https://api.github.com/users/" + username)
	if err != nil {
		return fmt.Errorf("failed to check GitHub username: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("github username does not exist")
	}

	return nil
}
