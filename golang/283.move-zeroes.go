/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int)  {
	/**
		key: swap the nonzero in front
	*/
	check:=0
	for i:=0; i<len(nums); i++ {
		if nums[i] !=  0 {
			// swap i to check
			nums[check], nums[i] = nums[i], nums[check]
			check++
		}
	}
}
// @lc code=end
