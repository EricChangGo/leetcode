#
# @lc app=leetcode id=1 lang=python3
#
# [1] Two Sum
#

# @lc code=start
class Solution:
    
    ### find left first -> find it in search_by_index
    def twoSum(self, nums, target):
        index_left = None
        index_dict = {}
        for index_right, num_right in enumerate(nums):
            num_left = target - num_right
            index_dict, index_left = self.search_by_index(index_dict, num_left, index_right, num_right)
            if index_left is not None and index_left is not index_right:
                return [index_left,index_right]
    def search_by_index(self, index_dict, target_num, index, num):
        target_index = index_dict.get(target_num, None)
        if target_index is None:
            index_dict[num] = index
            return index_dict, None
        else:
            return index_dict, target_index
# @lc code=end

