package leetcode

// @leet start
func lengthOfLastWord(s string) int {
	// Approach 1: using std libraries
	// Trim the surrounding whitespace and then split on whitespace.
	// Return the length of the last word in the array.

	// words := strings.Split(strings.TrimSpace(s), " ")
	// return len(words[len(words)-1])

	// Approach 2: using no libraries
	// Iterate from the end the string. Start counting at the first non
	// whitespace character and finish counting at the next whitespace
	// character.

	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			length++
		} else if length > 0 {
			return length
		}
	}
	return length
}

// @leet end
