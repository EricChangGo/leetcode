/*
 * @lc app=leetcode id=1200 lang=golang
 *
 * [1200] Minimum Absolute Difference
 */

// @lc code=start
func minimumAbsDifference(arr []int) [][]int {
	/*
			Two method
			1. O(n^2) to find all the minimum different value pairs and O(klogk) to sort the pairs 
				 -> worst O(n^2)
				 
			2. O(nlogn) Sort the arr, then with O(n) to compare every numbers with its own right adjacent member  -> worst O(nlogn)
	*/
	if len(arr) < 1 {
			return nil
	}
	
	sort.Ints(arr)
	diff:= arr[1] - arr[0]
	temp_diff:= 0
	pairs:= [][]int{}
	
	for i:=0; i<len(arr)-1; i++ {
			temp_diff = arr[i+1] - arr[i]
			if temp_diff < diff {
					diff = temp_diff
					pairs = nil
			}
			if temp_diff == diff {
					pairs = append(pairs, []int{arr[i], arr[i+1]})
			}
	}
	
	return pairs
}

// @lc code=end

