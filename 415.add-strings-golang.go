package leetcode

// @leet start
func addStrings(num1 string, num2 string) string {
	ret, carry := "", 0

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		if i >= 0 {
			carry += int(num1[i] - '0')
		}
		if j >= 0 {
			carry += int(num2[j] - '0')
		}
		ret = string(rune(carry%10)+'0') + ret
		carry /= 10
	}

	return ret
}

// @leet end
