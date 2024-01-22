package leetcode

// @leet start
func longestConsecutive(nums []int) int {
	// Approach:
	// Create a hashmap that maps number in nums -> int pointer to the length of
	// the sequence involving the number seen in nums. All numbers in a sequence
	// in the hashmap should point to the same length. Iterate over nums to
	// populate the hashmap. Iterate over the hashmap's values to find the max
	// length of sequences.

	// Time complexity: O(n)
	// Space complexity: O(n)

	// sequences := make(map[int]*int)
	//
	// for _, num := range nums {
	// 	if length, exists := sequences[num]; !exists {
	// 		// Initialize length for this new sequence.
	// 		sequences[num] = new(int)
	// 		*sequences[num] = 1
	// 		lengthLeft, leftExists := sequences[num-1]
	// 		if leftExists {
	// 			*length += *lengthLeft
	// 			lengthLeft = length
	// 		}
	// 		lengthRight, rightExists := sequences[num+1]
	// 		if rightExists {
	// 			*length += *lengthRight
	// 			lengthRight = length
	// 		}
	// 	}
	// }
	//
	// maxLength := 0
	// for _, length := range sequences {
	// 	if *length > maxLength {
	// 		maxLength = *length
	// 	}
	// }
	//
	// return maxLength

	// This is trash. I think I had an interesting idea, might be similar to a quick union, but executed poorly and
	// not as efficient. Maybe I'll revisit later.

	// Approach:
	// Add all numbers in nums to a set. Iterate over nums and check if the
	// current num is the start of a sequence (if the num - 1 exists in the set).
	// If num is the start of a sequence, then check if the next value in the
	// sequence exists. If it does then increment the length of the sequence and
	// continue checking until the sequence breaks. Keep track of the longest
	// sequence and return it.

	set := make(map[int]struct{}, len(nums))
	var longestSeq int

	for _, num := range nums {
		set[num] = struct{}{}
	}

	for num := range set {
		if _, exists := set[num-1]; !exists {
			length := 1
			for _, exists := set[num+length]; exists; _, exists = set[num+length] {
				length++
			}
			if length > longestSeq {
				longestSeq = length
			}
		}
	}

	return longestSeq
}

// @leet end

