/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
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
    Merge sort for O(NlogN) time complexity.
    The Difficulty is how to decide the divide the link this into two parts.
*/

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    /**
        end flag goes twice faster than mid flag
        when end gets to the end point, the mid gets to the mid point
    */
    mid, end := head, head
    for end.Next != nil && end.Next.Next != nil {
        mid = mid.Next
        end = end.Next.Next
    }
    
    // split the link
    temp := mid
    mid = mid.Next
    temp.Next = nil
    
    left, right := sortList(head), sortList(mid)
    return merge(left, right)
}

func merge(left *ListNode, right *ListNode) *ListNode{
    newHead := &ListNode{}
	// Set ptr to this new head node
	var linker *ListNode
    linker = newHead 
    
    // Start to link left and right into new link list
    // Use the newHead node's -> next as the start of the link list
    for left != nil && right != nil {
        if left.Val < right.Val {
            // use next as temp
            linker.Next = left
            left = left.Next             
        } else {
            linker.Next = right
            right = right.Next 
        }
        linker = linker.Next
    }
    
    // Left overs
    if left != nil {
		linker.Next = left
	} else if right != nil {
		linker.Next = right
	}
    return newHead.Next
}
// @lc code=end
