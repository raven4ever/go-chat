package utils

// IsValidPortNumber checks if a given integer value is a valid port number.
// A valid port number is between 1024 and 65535 (inclusive).
func IsValidPortNumber(value int) bool {
	return value > 1023 && value < 65536
}
