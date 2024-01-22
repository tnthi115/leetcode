package leetcode

// @leet start
func topKFrequent(nums []int, k int) []int {
	// Approach 1:
	// Use a min heap. Add all elements of nums to a min heap of size k.
	// Afterwards, the largest k elements will remain in the min heap because
	// once the min heap is full, each push of a new element will cause the
	// smallest element to be popped, leaving the largest k elements of nums.
	// This will have time complexity of O(n log n), where n is the size of nums.
	// You can use a max heap as well. Add all elements of nums to a max heap of size n, and
	// then pop off k elements. This has time complexity of O(n log n + k)? Instead, you can
	// heapify (turn a list into a heap in-place) in O(n) time, and then pop off k elements
	// in O(k log n) time. We can actually do better.

	// Approach 2;
	// Neetcode video: https://youtu.be/YPTqKIgVk-k
	// Bucket sort. Count all occurences of each element with a hashmap. Then map
	// count/frequency to list of elements that have that count/frequency in a
	// bounded array of length n (or n + 1 if we ignore index 0) where the
	// count/frequency is mapped to the indices of the array. We know that this
	// array's length can be bounded by n because in the worst case, nums can
	// have n of the same element, causing the count to be n. Finally, we can
	// iterate through this array starting at the end and add k elements to our
	// output array.
	// Time complexity: O(n + n + k) = O(n); n >= k
	// Space complexity: O(n + n) = O(n)

	count := make(map[int]int, len(nums))
	buckets := make([][]int, len(nums)+1)
	for _, n := range nums {
		count[n]++
	}
	for n, c := range count {
		buckets[c] = append(buckets[c], n)
	}
	topK := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, n := range buckets[i] {
			topK = append(topK, n)
			if len(topK) == k {
				return topK
			}
		}
	}
	// This won't be reached.
	return topK
}

// @leet end

