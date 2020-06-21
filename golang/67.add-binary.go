/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	/**
		Solutions
		1. transfer to int and use bitOR ^ -> have limitation
		2. [Better] string maping -> augend(longer) + addend(short)
			key: int to char, char to int
	*/
	var aStr *string
	var bStr *string
	aStr = &a
	bStr = &b
	countA := len((*aStr))
	countB := len((*bStr))
	
	carry := 0
	tempSum := 0
	totalStr := ""
	for countA > 0 || countB > 0 {
		tempSum = carry
		if countA > 0 {
            tempSum += (int((*aStr)[countA-1]) - '0')
			countA --
		}
		if countB > 0 {
            tempSum += (int((*bStr)[countB-1]) - '0')
			countB --
		}
		carry = tempSum / 2
        totalStr = addTotalStr(tempSum % 2, totalStr)
	}
    
    if carry == 1 {
		totalStr = addTotalStr(carry, totalStr)
    }
    
	return totalStr
}

func addTotalStr(carry int, totalStr string) string {
    return strconv.Itoa(carry) + totalStr
}
// @lc code=end
