/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	if n <= 2 {
			return n
	}
	
	steps := []int{1, 2} // n = 1 & 2
	current := 0
	for i:=3; i<=n; i++ {
			current = steps[0] + steps[1]
			steps[0] = steps[1]
			steps[1] = current
			// dp -> steps[i] = steps[i-1] + steps[i-2]
	}
	return current
}
// @lc code=end

