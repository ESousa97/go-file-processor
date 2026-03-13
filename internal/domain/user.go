package domain

// User represents a user entity in the system.
// It is the primary data structure used for CSV mapping and JSON serialization.
type User struct {
	// ID is the unique identifier for the user.
	ID string `json:"id"`
	// Name is the full name of the user.
	Name string `json:"name"`
	// Email is the primary contact email for the user.
	Email string `json:"email"`
	// Role defines the access level or position of the user within the system.
	Role string `json:"role"`
}
