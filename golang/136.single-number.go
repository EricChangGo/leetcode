/**
*	Goal: Implement it without using extra memory
*	The key is to use bitwise XOR ^ operator, which gains Commutative Property.
*	[2,3,1,2,3]
*   - 2^3^1^2^3 = 2^2^3^3^1
**/
func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
