/*
 * @lc app=leetcode id=912 lang=cpp
 *
 * [912] Sort an Array
 */

// @lc code=start
class Solution {
// practice quick sort
public:
    vector<int> sortArray(vector<int>& nums) {
      quickSort(nums, 0, nums.size()-1);
      return nums;
    }
    
    void quickSort(vector<int>& nums, int start_i, int end_i) {
      if(start_i < end_i) {
        int pivot = partition(nums, start_i, end_i);
        quickSort(nums, start_i, pivot-1);
        quickSort(nums, pivot+1, end_i);  
      }
    }
    
    int partition(vector<int>& nums, int start_i, int end_i) {
      int pivot = nums[end_i];
      int small_end = start_i;
      
      for(int large_end = start_i; large_end <= end_i; large_end++) {
        if(pivot > nums[large_end]) {
          swap(&nums[small_end], &nums[large_end]);
          small_end ++;
        }
      }
      
      // pivot is still in the end,
      // swap it with the first large number, turn it into the boarder of the two area
      swap(&nums[end_i], &nums[small_end]);
      return small_end;
    }
    
    void swap(int* a, int* b) {
      int temp = *a;
      *a = *b;
      *b = temp;
    }
};
// @lc code=end

