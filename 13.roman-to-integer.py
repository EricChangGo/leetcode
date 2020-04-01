#
# @lc app=leetcode id=13 lang=python3
#
# [13] Roman to Integer
#

# @lc code=start
class Solution:
    symbol_dict = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000
    }
                
    def symbol_to_value(self, sym):
        return self.symbol_dict[sym]
    
    def romanToInt(self, s: str) -> int:
        ### label all the char
        # default maxDigit = s[0] value
        temp_value = self.symbol_to_value(s[0])
        total_value = 0
        for i in range(1, len(s)):
            value = self.symbol_to_value(s[i])
            check_value = self.symbol_to_value(s[i-1])
            if value > check_value:
                # Larger than previous
                temp_value = value - temp_value
            elif value < check_value:
                # Smaller than previous -> previous add into total, set temp = new value
                total_value += temp_value
                temp_value = value
            elif value == check_value:
                temp_value += value

        total_value += temp_value
        return total_value
# @lc code=end

