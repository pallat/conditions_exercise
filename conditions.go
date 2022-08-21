package conditions

// isOdd returns true if the number is odd.
// % operator is used to check if the remainder of the division is 0.
func isOdd(n int) bool {
	return n%2 == 1
}
