# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a CLI tool that fetches and displays GitHub user activity using the GitHub API. The project follows the roadmap.sh GitHub User Activity project specifications (https://roadmap.sh/projects/github-user-activity).

**Current Status**: This is a new project with only the go.mod file initialized. The main implementation is yet to be built.

## Project Requirements

The CLI should:
- Accept a GitHub username as argument: `github-activity <username>`
- Use GitHub API endpoint: `https://api.github.com/users/<username>/events`
- Display recent activity in the terminal (e.g., "Pushed 3 commits to repo", "Opened a new issue")
- Handle errors gracefully (invalid usernames, API failures)
- Use only standard library (no external dependencies for API calls)

## Development Commands

Since this is a new Go project, use these standard Go commands:

```bash
# Initialize and build
go mod tidy
go build -o github-activity .

# Run the application
go run main.go <username>
# Or after building:
./github-activity <username>

# Test the application
go test ./...

# Format code
go fmt ./...

# Vet code for issues
go vet ./...
```

## Architecture Notes

This is a simple CLI application that should:
- Have a main.go file as the entry point
- Parse command line arguments for the username
- Make HTTP requests to GitHub API using standard net/http
- Parse JSON response and format output for display
- Implement proper error handling for network and API errors

## Module Information

- Module: `github.com/jeremycardona/github-user-activity`
- Go Version: 1.25.0
- Dependencies: Should use only standard library as per project requirements

## Do not
- do not implement any code.
- do not write any code.