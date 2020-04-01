#
# @lc app=leetcode id=9 lang=python3
#
# [9] Palindrome Number
#

# @lc code=start
class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False

        # set i into str -> able to recognize the maximum number of digits
        checkX = str(x)
        maxLenX = len(checkX)
        for i in range(maxLenX):
            if checkX[i] != checkX[maxLenX-1-i]:
                return False
        return True
# @lc code=end

