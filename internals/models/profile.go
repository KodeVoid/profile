package models

// User represents a user profile
type User struct {
	Email string `json:"email"`

	Name  string `json:"name"`
	Stack string `json:"stack"`
}

// Profile represents the full profile response
type Profile struct {
	Status string `json:"status"`
	User   User   `json:"user"`

	Timestamp string `json:"timestamp"`
	Fact      string `json:"fact"`
}
