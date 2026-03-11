package domain

// User represents a user entity in the system.
// This is our target domain model for the CSV transformation.
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
