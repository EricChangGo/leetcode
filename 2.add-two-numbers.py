# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def addTwoNumbers(self, l1, l2):
        carry = None
        temp_digit = None
        ans_link = None
        temp_link = None
        while True:   
            l1_num = 0
            l2_num = 0
            if l1:
                l1_num = l1.val
                l1 = l1.next
            if l2:
                l2_num = l2.val
                l2 = l2.next
            
            temp_digit = l1_num + l2_num
            if carry:
                temp_digit += carry
            carry, temp_digit = set_carry(temp_digit)
            #print(temp_digit)
            temp_node = ListNode(temp_digit)
            #print(temp_node.val)
            
            if temp_link is None:
                temp_link = temp_node
                ans_link = temp_link
            else:
                temp_link.next = temp_node
                temp_link = temp_link.next
            #print(temp_link.val)
            if l1 is None and l2 is None and carry is 0:
                break
        return ans_link
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
def set_carry(num):
    if num < 10:
        return 0, num
    return int(num / 10), num % 10
# @lc code=end

