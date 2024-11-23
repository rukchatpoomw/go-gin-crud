package utils

// CalculateSkip calculates the skip value for pagination
func CalculateSkip(page int64, limit int64) int64 {
	return (page - 1) * limit
}
