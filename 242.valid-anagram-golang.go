package leetcode

import (
	"slices"
)

// @leet start
func isAnagram(s string, t string) bool {
	// Approach:
	// Create a counts map that keeps track of the counts of each letter in
	// string s. Count the letters in s, then decrement the counts for each
	// letter in t. If the strings are anagrams, all counts in the end should be
	// 0.

	// Time complexity: O(N + M + N) = O(N + M)
	// Space complexity: O(N)

	counts := make(map[rune]int)

	// Count letters in s
	for _, letter := range s {
		counts[letter]++
	}

	// Decrement counts for letters in t
	for _, letter := range t {
		counts[letter]--
		if counts[letter] < 0 {
			// If the letter is not present in s or the count becomes negative,
			// the strings are not anagrams.
			return false
		}
	}

	// Check if all counts are 0
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

func _isAnagram(s string, t string) bool {
	// Approach:
	// Sort each string and compare the sorted strings. If the strings are
	// anagrams, the sorted strings will be equal.

	// Time complexity: O(N log N)
	// Explanation: The sorting algorithm in Go has time complexity of O(n log n).

	// Space complexity: O(log (max(N, M)))
	// Explanation: The sorting algorithm in Go has space complexity O(log n)
	// where n is the size of the input.

	sortedS := sortString(s)
	sortedT := sortString(t)
	return sortedS == sortedT
}

func sortString(s string) string {
	chars := []rune(s)
	slices.Sort(chars)
	return string(chars)
}

// @leet end

