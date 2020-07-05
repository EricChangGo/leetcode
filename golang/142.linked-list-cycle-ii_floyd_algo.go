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
 func hasCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
        slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
            // has cycle
            /** 
                Important part:
                
                The two will met in point of the beginning point of cycle minus N steps;
                N will be the the steps from the beginning of the link list to the 
                begin of the cycle.
                
                Math prove: https://www.youtube.com/watch?v=LUm2ABqAs1w
            */
			return slow
		}
	}
    // no cycle
    return nil
}

func detectCycle(head *ListNode) *ListNode {
    metNode := hasCycle(head)
    if metNode == nil {
        return nil
    }
    
    chaseNode := head    
    for metNode != chaseNode {
        chaseNode = chaseNode.Next
        metNode = metNode.Next
    }
    return chaseNode
}
// @lc code=end
