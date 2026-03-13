package processor

import (
	"go-file-processor/internal/domain"
	"regexp"
)

// Transformer defines a function that modifies a User or returns false to filter it out.
// Following the Middleware/Chain of Responsibility pattern.
type Transformer func(user *domain.User) bool

// EmailFilter returns a transformer that filters users with invalid emails based on regex.
func EmailFilter(pattern string) Transformer {
	re := regexp.MustCompile(pattern)
	return func(user *domain.User) bool {
		return re.MatchString(user.Email)
	}
}

// FieldMasker returns a transformer that masks specific fields.
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

// RoleFilter returns a transformer that only allows specific roles.
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
