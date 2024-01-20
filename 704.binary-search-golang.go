package leetcode

// @leet start
func search(nums []int, target int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		// Using left + (right - left) / 2 might be better than (left + right) / 2
		// since left + right might cause an overflow issue if both left and right
		// are very large.
		mid := left + (right-left)/2
		val := nums[mid]
		if val == target {
			return mid
		} else if val < target {
			left = mid + 1
		} else /* if val > target */ {
			right = mid - 1
		}
	}
	return -1
}

// @leet end
