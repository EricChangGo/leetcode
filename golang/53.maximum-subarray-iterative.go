/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	/*
		O(N) Solution
	*/
	MaxInt:= (1 << 31)-1
	sum:= -MaxInt-1
	temp_sum:= 0
	
	for i:=0; i<len(nums); i++ {
		if temp_sum < 0 {
			// drop previous numbers
			temp_sum = 0
		}
		temp_sum += nums[i]
		if temp_sum > sum {
			// get better result
			sum = temp_sum
		}
	}
	return sum
}
// @lc code=end

