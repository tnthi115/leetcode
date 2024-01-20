package leetcode

// Recursive approach
func fib_recursive(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// @leet start
// Iterative approach
func fib(n int) int {
	if n <= 1 {
		return n
	}
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

// @leet end
