package leetcode

import "slices"

// @leet start
func _groupAnagrams(strs []string) [][]string {
	// Approach:
	// Iterate over strs and add each string to a map(sorted string -> list of
	// strings that are anagrams). We are using sorting as a hash function on the
	// strings. We can replace sorting with any more efficient hash function. See
	// the below approach for a more optimal solution. After building the map, we
	// can iterate over it and generate our output list.

	// Time complexity: O(m * nlogn) where m is len(strs) and n is average length
	// of each string in strs

	buckets := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		sortedStr := string(runes)
		buckets[sortedStr] = append(buckets[sortedStr], str)
	}

	groupedAnagrams := make([][]string, 0, len(buckets))
	for _, bucket := range buckets {
		groupedAnagrams = append(groupedAnagrams, bucket)
	}

	return groupedAnagrams
}

func groupAnagrams(strs []string) [][]string {
	// Approach:
	// Iterate over strs and add each string to a map(hashed string -> list of
	// strings that are anagrams). Each string consists of lowercase English
	// letters, so we know we are limited to 26 characters. We can create a
	// simple hashing function that counts each letter in a string and stores the
	// counts in a slice. After building the map, we can iterate over it and
	// generate our output list.

	// Time complexity: O(m * n * 26) = O(m * n) where m is len(strs) and n is
	// the average length of each string in strs of each string in strs

	buckets := make(map[[26]int][]string)
	for _, str := range strs {
		var countKey [26]int
		for _, char := range str {
			countKey[char-'a']++
		}
		buckets[countKey] = append(buckets[countKey], str)
	}

	groupedAnagrams := make([][]string, 0, len(buckets))
	for _, bucket := range buckets {
		groupedAnagrams = append(groupedAnagrams, bucket)
	}

	return groupedAnagrams
}

// @leet end

