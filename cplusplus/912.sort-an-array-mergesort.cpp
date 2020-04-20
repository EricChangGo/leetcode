/*
 * @lc app=leetcode id=912 lang=cpp
 *
 * [912] Sort an Array
 */

// @lc code=start
class Solution {
// practice merge sort
public:
    vector<int> sortArray(vector<int>& nums) {
      merge_sort(nums, 0, nums.size()-1);
      return nums;
    }
    
    void merge_sort(vector<int>& nums, int start, int end) {
      // stop when the we target the basic one element
      if(start >= end) return; 
      
      // divide the array into half
      int mid = (start+end)/2;
      // sort left elements
      merge_sort(nums, start, mid);
      // sort right elements
      merge_sort(nums, mid+1, end);
      
      merge(nums, start, mid, end);
    }
    
    void merge(vector<int>& nums, int start, int mid, int end) {
      /*
        1. Set left <= mid, right > mid
        2. Arrange into a new array with all left right element
        
        Ex: left = [1,5], right = [3,8]
        New array = [1,3,5,8]
      */
      vector<int> left(nums.begin()+start, nums.begin()+mid+1);
      vector<int> right(nums.begin()+mid+1, nums.begin()+end+1);
      left.push_back(INT_MAX);
      right.push_back(INT_MAX);
      
      int l_idx = 0;
      int r_idx = 0;
      
      for(int i=start; i<=end; i++) {
        // set nums[i] 
        if(left[l_idx] <= right[r_idx]) {
          nums[i] = left[l_idx];
          l_idx++;
        }
        else {
          nums[i] = right[r_idx];
          r_idx++;
        }
      }
    }
  
    void print_list(vector<int>& list) {
      for(const auto &it: list) {
        cout << it << " ";
      }
      cout << endl;
    }
};
// @lc code=end

