package leetcode

import (
	"strings"
	"unicode"
)

// @leet start
func isPalindrome(s string) bool {
	// Approach:
	// 1. Remove all nonalphanumeric characters in s.
	// 2. Create two pointers, one starting at the beginning of s and the other at the end. Iterate inward
	// until the pointers move past each other, checking that each character at the two pointers match.

	// Time complexity: O(n)
	// Space complexity: O(1)

	s = strings.Map(removeNonAlphanumeric, s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// removeNonAlphanumeric is a mapping function for strings.Map to remove
// non-alphanumeric characters (non-letter and non-number)
func removeNonAlphanumeric(r rune) rune {
	if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
		return -1
	}
	return unicode.ToLower(r)
}

// @leet end

