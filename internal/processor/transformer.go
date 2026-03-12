package processor

import (
	"go-file-processor/internal/domain"
	"regexp"
)

// Transformer defines a function that modifies a User or returns false to filter it out.
// Following the Middleware/Chain of Responsibility pattern.
type Transformer func(user *domain.User) bool

// EmailValidatorTransformer returns a transformer that validates email format.
func EmailValidatorTransformer() Transformer {
	// Simple email regex for demonstration
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return func(user *domain.User) bool {
		return re.MatchString(user.Email)
	}
}

// SensitiveDataTransformer returns a transformer that masks specific fields.
func SensitiveDataTransformer(maskRole bool) Transformer {
	return func(user *domain.User) bool {
		if maskRole {
			user.Role = "CONFIDENTIAL"
		}
		return true
	}
}
