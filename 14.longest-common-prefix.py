#
# @lc app=leetcode id=14 lang=python3
#
# [14] Longest Common Prefix
#

# @lc code=start
class Solution:
    def check_prefix(self, on_check_arr: list, check_char: str, index_flag: int):
        for word in on_check_arr:
            if word[index_flag] != check_char:
                return False
        return True
    
    def get_min_word_size(self, on_check_arr: list, min_word_size: int):
        for word in on_check_arr:
            word_len = len(word)
            if word_len < min_word_size:
                min_word_size = word_len
        return min_word_size
    
    def longestCommonPrefix(self, strs: List[str]) -> str:
        # get the smallest word to limit the checking
        if not strs or not strs[0]:
            return ""

        min_word_size = len(strs[0])
        min_word_size = self.get_min_word_size(strs[1::], min_word_size)
        
        index_flag = 0
        has_prefix = True
        check_char = strs[0][index_flag]
        len_check_word = len(strs[0])

        while(1):
            if index_flag == min_word_size:
                break
            check_char = strs[0][index_flag]
            has_prefix = self.check_prefix(strs[1::], check_char, index_flag)

            if not has_prefix:
                # not match, break while
                break
            index_flag += 1

        if index_flag == 0:
            if min_word_size == 0 or not has_prefix:
                # contain "" in list, or no prefix at all
                return ""
            return check_char
        return strs[0][:index_flag]
# @lc code=end

