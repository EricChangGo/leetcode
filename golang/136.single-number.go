/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
// Goal: Implement it without using extra memory

func singleNumber(nums []int) int {
	var res int
	var noFound bool
	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			continue
		}
		res = nums[i]
		noFound = true
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				nums[i] = -1
				nums[j] = -1
				noFound = false
				break
			}
		}

		if noFound {
			return res
		}
	}
	return -1
}

// @lc code=end
