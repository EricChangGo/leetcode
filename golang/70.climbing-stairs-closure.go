/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
/** using function closure*/
func fibonacci() func() int {
	x, y := 1, 2
	return func() int {
		x, y = y, x + y
		return y
	}
}a

func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    
    f := fibonacci()
    current:= 0
    for i:=3; i<=n; i++ {
        current = f()
    }
    return current
}
// @lc code=end

