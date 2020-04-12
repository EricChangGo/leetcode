/*
 * @lc app=leetcode id=33 lang=javascript
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 
 Binary Search O(logn) by sorted array
 */

var search = function(nums, target) { 
  return binarySearch(nums, 0, nums.length-1, target);
};

var binarySearch = (nums, start_index, end_index, target) => {
  
  // search for target
  let start = nums[start_index];
  let end = nums[end_index];
  let middle_index = Math.floor((start_index+end_index)/2);
  let middle = nums[middle_index];
  
  if(target === end){
    return end_index;
  }
  else if(target === start){
    return start_index;
  }
  else if(target === middle) {
    return middle_index;
  }
  if (start_index >= end_index) return -1;
 
  if(middle > target) {
    // left, turning point is in the left
    // target is between turning point and end point
    if(middle > end && end > target) return binarySearch(nums, middle_index+1, end_index, target);
    // right, substain normal binary search
    else return binarySearch(nums, start_index, middle_index-1, target);
  }
  else {
    // right, turning point is in the right
    // target is between turning point and 
    if(middle < start && start < target) return binarySearch(nums, start_index, middle_index-1, target); 
    // left, substain normal binary search
    else return binarySearch(nums, middle_index+1, end_index, target);
  }
};
// @lc code=end

