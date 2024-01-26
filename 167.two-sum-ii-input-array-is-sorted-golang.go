package leetcode

// @leet start
func twoSum(numbers []int, target int) []int {
	// Approach:
	// Use two pointers, one at the beginning of the array and one at the end.
	// Calculate the sum of numbers[l] + numbers[r]. If the sum equals target,
	// return [l + 1, r + 1]. If the sum is less than target, increment l to try
	// a higher value. If the sum is greater than target, decrement r to try a
	// lower value.

	for l, r := 0, len(numbers)-1; l < r; {
		sum := numbers[l] + numbers[r]
		switch {
		case sum > target:
			r--
		case sum < target:
			l++
		default:
			return []int{l + 1, r + 1}
		}
	}

	return []int{-1, -1} // Unreachable in this specific problem context
}

// @leet end
