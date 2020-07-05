/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 /**
 	Using Map - cost memory space will be much larger
 */
 func detectCycle(head *ListNode) *ListNode {
    var startNode *ListNode
    nodeMap := make(map[*ListNode]int)
    for head != nil {
        if _, ok := nodeMap[head]; ok {
            return head
        }
        nodeMap[head] = 0
        head = head.Next
    }
    return startNode
}
// @lc code=end
