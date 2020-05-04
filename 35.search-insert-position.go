/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	// head, tail 
	start := 0
	end := len(nums)-1
	insert_idx := -1
	if target < nums[0] {
			return 0
	}
	if target > nums[end] {
			return end+1
	}
	
	for insert_idx < 0 {
			check := (start+end)/2
			
			if check == start {
					// find insert point, insert between the start end
					insert_idx = end
					if target <= nums[start] {
							 // target value equals to start, insert in front of start
							 insert_idx = start
					}
			}
			
			if target < nums[check] {
					end = check
			} else { 
					start = check
			}
	}
	return insert_idx
}
// @lc code=end

