/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
const MaxInt = (1 << 31) - 1
const MinInt = -MaxInt - 1

func max(a,b int) int {
	if b > a {
		return b
	}
	return a
}

func getLeftMax(num []int, start int, end int) int {
	temp:=0
	maxValue:=MinInt
	for i:=end; i>=start; i-- {
		temp+=num[i]
		maxValue=max(maxValue, temp)
	}
	return maxValue
}

func getRightMax(num []int, start int, end int) int {
	temp:=0
	maxValue:=MinInt
	for i:=start; i<=end; i++ {
		temp+=num[i]
		maxValue=max(maxValue, temp)
	}
	return maxValue
}

func maxSubArray(nums []int) int {
	return getMaxSubs(nums, 0, len(nums)-1)
}

func getMaxSubs(nums []int, start int, end int) int {
	if start == end {
		return nums[start]
	}
	mid := (start+end)/2
	leftSub := getMaxSubs(nums, start,  mid)
	rightSub := getMaxSubs(nums, mid+1, end)

	leftMax := getLeftMax(nums, start, mid)
	rightMax := getRightMax(nums, mid+1, end)
	total := leftMax + rightMax
	return max(max(leftSub, rightSub), total)
}
// @lc code=end

