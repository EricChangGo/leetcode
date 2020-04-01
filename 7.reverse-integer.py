#
# @lc app=leetcode id=7 lang=python3
#
# [7] Reverse Integer
#

# @lc code=start
class Solution:
    def reverse(self, x):
        string_input = str(x)

        if  49 <= ord(string_input[0]) and ord(string_input[0]) <= 57:
            # first char is number 
            string_input = string_input[::-1]
            if string_input[0] is '0':
                string_input = string_input[1::]
        else:
            numbers = string_input[1::]
            numbers = numbers[::-1]
            string_input = '-{}'.format(numbers)
        
        try:
            ans_int = int(string_input)
            if ans_int > 2147483647 or ans_int < -2147483648:
                 return 0
        except:
            return 0
            
        return ans_int
# @lc code=end

