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
	
	ways := make([]int, n+1, n+1)
	ways[1] = 1
	ways[2] = 2
	for i:=3; i<=n; i++ {
			ways[i] = ways[i-1] + ways[i-2]
	}
	return ways[n]
}
// @lc code=end

