# GitHub User Activity CLI - Development Todo

## Phase 1: Project Structure Setup
- [x] Create todo.md file
- [x] Create main.go with basic structure
- [ ] Set up basic project organization

## Phase 2: Required Functions

### Command Line Interface Functions
- [ ] `parseArgs()` - Parse and validate command line arguments
  - Check if username provided
  - Handle help flags (-h, --help)
  - Return username or error
- [ ] `printUsage()` - Display help/usage information
  - Show correct command format
  - Display example usage
- [ ] `validateUsername()` - Basic username validation
  - Check for empty/whitespace
  - Basic format validation

### GitHub API Integration Functions
- [ ] `buildAPIURL(username string) string` - Construct GitHub API endpoint
  - Format: `https://api.github.com/users/{username}/events`
- [ ] `fetchUserEvents(username string) ([]byte, error)` - Make HTTP request
  - Create HTTP client
  - Make GET request to API
  - Return response body or error
- [ ] `handleAPIResponse(resp *http.Response) error` - Process HTTP response
  - Check status codes (200, 404, 403, etc.)
  - Handle rate limiting
  - Return appropriate errors

### Data Processing Functions
- [ ] `parseEventsJSON(data []byte) ([]GitHubEvent, error)` - Parse JSON response
  - Unmarshal JSON into structs
  - Handle parsing errors
- [ ] `formatEvent(event GitHubEvent) string` - Format single event for display
  - Handle different event types (PushEvent, IssuesEvent, etc.)
  - Return human-readable string
- [ ] `displayActivity(events []GitHubEvent)` - Output formatted activity
  - Loop through events
  - Print formatted strings

### Error Handling Functions
- [ ] `handleNetworkError(err error) string` - Handle network/HTTP errors
  - Connection timeouts
  - DNS resolution errors
- [ ] `handleAPIError(statusCode int) string` - Handle GitHub API errors
  - 404: User not found
  - 403: Rate limit exceeded
  - 422: Invalid request
- [ ] `handleJSONError(err error) string` - Handle JSON parsing errors
  - Invalid JSON format
  - Missing required fields

## Phase 3: Data Structures

### Core Structs
- [ ] `GitHubEvent` - Main event structure
  - Type (string)
  - Actor (User info)
  - Repo (Repository info)
  - Payload (interface{} for different event types)
  - CreatedAt (time)
- [ ] `Actor` - User information
  - Login (string)
  - ID (int)
- [ ] `Repository` - Repository information
  - Name (string)
  - URL (string)

### Event-Specific Payload Structs
- [ ] `PushEventPayload` - For push events
  - Commits ([]Commit)
  - Size (int)
- [ ] `IssuesEventPayload` - For issue events
  - Action (string)
  - Issue (Issue struct)
- [ ] `PullRequestEventPayload` - For PR events
  - Action (string)
  - PullRequest (PR struct)

## Phase 4: Implementation Steps
- [ ] Start with main.go and basic CLI parsing
- [ ] Implement HTTP client and test with simple API call
- [ ] Add JSON structs and parsing
- [ ] Implement event type detection and formatting
- [ ] Add comprehensive error handling
- [ ] Test with various usernames and scenarios

## Phase 5: Testing & Polish
- [ ] Test with valid GitHub usernames
- [ ] Test error scenarios:
  - [ ] Invalid/non-existent username
  - [ ] Network connectivity issues
  - [ ] Rate limiting
- [ ] Run `go fmt ./...`
- [ ] Run `go vet ./...`
- [ ] Build final binary: `go build -o github-activity .`
- [ ] Test built binary

## Learning Objectives
- Understanding Go's net/http package for HTTP clients
- JSON parsing with encoding/json package
- Command-line argument handling with os.Args
- Error handling patterns in Go
- Struct composition and interfaces
- Working with GitHub's REST API