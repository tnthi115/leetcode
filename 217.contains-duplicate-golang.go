package leetcode

// @leet start
func containsDuplicate(nums []int) bool {
	// This approach has time complexity O(n log n) because and space complexity O(log n) because of sorting.
	// slices.Sort(nums)
	// for i := 1; i < len(nums); i++ {
	// 	if nums[i-1] == nums[i] {
	// 		return true
	// 	}
	// }
	// return false

	// This approach has time complexity O(N) but space complexity O(N) due to the hashmap.
	// This is how to create a seen in Go idomatically. We use the empty struct as
	// the type for the values of the map because the values associated with the
	// keys are irrelevant, and the empty struct struct{} has no size, which
	// saves memory.
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, duplicate := seen[num]; duplicate {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

// @leet end

