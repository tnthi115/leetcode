package leetcode

// @leet start
func twoSum(nums []int, target int) []int {
	diffs := make(map[int]int)
	for i, val := range nums {
		diff := target - val
		if seen_diff, exists := diffs[diff]; exists {
			return []int{i, seen_diff}
		}
		diffs[val] = i
	}
	return []int{}
}

// @leet end
