package processor

import (
	"go-file-processor/internal/domain"
	"regexp"
)

// Transformer defines a function signature for modifying a [domain.User] record.
// It returns true if the record should be kept, or false to filter it out.
// This implements the Middleware/Chain of Responsibility pattern.
type Transformer func(user *domain.User) bool

// EmailFilter returns a [Transformer] that validates the user's email against a regex pattern.
// Records with emails that do not match the pattern are filtered out.
func EmailFilter(pattern string) Transformer {
	re := regexp.MustCompile(pattern)
	return func(user *domain.User) bool {
		return re.MatchString(user.Email)
	}
}

// FieldMasker returns a [Transformer] that obscures the content of specified fields
// for data privacy (e.g., masking emails or roles).
func FieldMasker(field string) Transformer {
	return func(user *domain.User) bool {
		switch field {
		case "email":
			user.Email = "****@****.***"
		case "role":
			user.Role = "CONFIDENTIAL"
		}
		return true
	}
}

// RoleFilter returns a [Transformer] that only allows users with roles
// specified in the allowedRoles slice.
func RoleFilter(allowedRoles []string) Transformer {
	roles := make(map[string]struct{})
	for _, r := range allowedRoles {
		roles[r] = struct{}{}
	}
	return func(user *domain.User) bool {
		_, ok := roles[user.Role]
		return ok
	}
}
