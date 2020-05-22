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

func getLeftMax(num []int, start int, end int) int {
	sum:=0
	maxValue:=-2199999
	for i:=end; i>=start; i-- {
		sum+=num[i]
		maxValue=max(maxValue, sum)
	}
	return maxValue
}

func getRightMax(num []int, start int, end int) int {
	sum:=0
	maxValue:=-2199999
	for i:=start; i<=end; i++ {
		sum+=num[i]
		maxValue=max(maxValue, sum)
	}
	return maxValue
}

func maxSubArray(nums []int) int {
	return getMaxSubs(nums, 0, len(nums)-1)
}

/**
	[-2, 1] -> 1
	[4, -1, 2] -> 4,-1,2
	Q. adding negative number or not?
*/
func getMaxSubs(nums []int, start int, end int) int {
	if start == end {
		return nums[start]
	}
	// divide the array into half
	mid := (start+end)/2
	left_sub := getMaxSubs(nums, start,  mid)
	right_sub := getMaxSubs(nums, mid+1, end)

	left_max := getLeftMax(nums, start, mid)
	right_max := getRightMax(nums, mid+1, end)
	total := left_max + right_max
	return max(max(left_sub, right_sub), total)
}
// @lc code=end

