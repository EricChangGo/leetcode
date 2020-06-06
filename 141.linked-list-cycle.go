/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	/**
		Set a key: value (*ListNode): ____
		1. node exist in map return true
		2. set node in map
	*/
	nodeMap := make(map[*ListNode]int)
	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return true
		}
		nodeMap[head] = 0
		head = head.Next
	}
	return false
}
// @lc code=end
