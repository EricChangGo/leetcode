/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func max(a,b int) int {
	if b > a {
		return b
	}
	return a
}

func maxSubArray(nums []int) int {
	MaxInt:= (1 << 31)-1
	MinInt:= -MaxInt-1

	temp:=0
    maxValue:=MinInt
    for i:=0; i<len(nums); i++ {
        if temp < 0 {
            temp = 0
        }
        temp+=nums[i]
        maxValue=max(maxValue, temp)
    }
    return maxValue
}
// @lc code=end

