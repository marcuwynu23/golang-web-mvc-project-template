package utils

// Ptr returns a pointer to the given value.
// Useful for tests or optional fields.
func Ptr[T any](v T) *T {
	return &v
}

