/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func binaryInsert(nums []int, len_nums int, target int) int {
	// head, tail
	start := 0
	end := len_nums-1
	insert_idx := -1
	check := 0
	for insert_idx < 0 {
			check = (start+end)/2
			
			if check == start {
					// find insert point, insert between the start end
					insert_idx = start+1
					if target == nums[start] {
							 // target value equals to start, insert in front of start
							 insert_idx -= 1
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

func linearInsert(nums []int, len_nums int, target int) int {
	for i:=0; i<len_nums-1; i++ {
			if nums[i] < target && target <= nums[i+1] {
					return i+1
			} 
	}
	return 0
}

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
			return 0
	}
	len_nums := len(nums)
	if target > nums[len_nums-1] {
			return len_nums
	}
	
	if(len_nums < 1000) {
			return linearInsert(nums, len_nums, target)
	}
	return binaryInsert(nums, len_nums, target) 
}
// @lc code=end

