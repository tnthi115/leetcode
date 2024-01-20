package leetcode

// @leet start
func plusOne(digits []int) []int {
	// Generic solution
	// carry := 1
	// for i := len(digits) - 1; i >= 0; i-- {
	// 	carry = carry + digits[i]
	// 	digits[i] = carry % 10
	// 	carry /= 10
	// }
	// if carry > 0 {
	// 	return append([]int{carry}, digits...)
	// }
	// return digits

	// Optimized solution for just adding 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			// We don't have a carry so we can just increment this digit
			// and return early.
			digits[i]++
			return digits
		}
		// We have a carry.
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

// @leet end

