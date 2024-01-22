package leetcode

// @leet start
func productExceptSelf(nums []int) []int {
	// Approach:
	// Create the answer array with the same length as nums. Iterate over the
	// nums array from start to end, keeping track of a rolling product of all
	// elements preceding the current one (prefix). Set answer[i] to the rolling
	// product (prefix). Do the same from end to start, multiplying the rolling
	// product (posfix) into answer[i] instead of setting (and overwriting).

	// Time complexity: O(n + n + n) = O(3n) = O(n)
	// Space complexity: O(1). The output array doesn't count as extra space for
	// space complexity analysis.

	res := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}

// @leet end

