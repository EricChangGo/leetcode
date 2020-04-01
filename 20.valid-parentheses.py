#
# @lc app=leetcode id=20 lang=python3
#
# [20] Valid Parentheses
#

# @lc code=start
class Solution:
    check_list = {
        '(': 1,
        '{': 2,
        '[': 3,
        
        ')': -1,
        '}': -2,
        ']': -3
    }
    
    def isValid(self, s: str) -> bool:
        if not s:
            return True
        
        len_str = len(s)
        if len_str % 2 > 0:
            # len_str is not even
            return False
        
        stack_for_check = [0]
        has_to_check = True
        
        if not self.check_list.get(s[0], None):
            # sign not in check_list
            return False

        for i in range(1, len_str):
            flag = self.check_list.get(s[i], None)
            if not flag:
                # sign not in check_list 
                return False
            if flag > 0:
                stack_for_check.append(i)
            else:
                check_sign_index = stack_for_check.pop()
                if self.check_list.get(s[check_sign_index]) + flag != 0:
                    # not match
                    return False

        if stack_for_check:
            return False
        return True
# @lc code=end

