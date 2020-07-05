/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
    Slow runner runs one step at a time
    Fast runner runs two step at a time
    
    1. no cycle -> the fast runner will reach the end point first
    2. cycle -> the fast runner will be n steps behind the second, and get close a step closer evertime -> eventually reach to the slow runner
*/

/** Using floyd algorithm */
func hasCycle(head *ListNode) bool {
    has_cycle := false
    if head == nil || head.Next == nil {
		return has_cycle
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
            has_cycle = true
			return has_cycle
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
    return has_cycle
}

/** Hash Table
func hasCycle(head *ListNode) bool {
    nodeMap := make(map[*ListNode]int)
    for head != nil {
        if _, ok := nodeMap[head]; ok {
            return true
        }
        nodeMap[head] = 0
        head = head.Next
    }
    return false
}*/