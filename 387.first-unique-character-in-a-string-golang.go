package leetcode

// @leet start
func firstUniqChar(s string) int {
	numLetters := 26
	counts := make([]int, numLetters)

	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if counts[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

// @leet end

