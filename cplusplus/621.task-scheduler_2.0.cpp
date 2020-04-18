/*
 * @lc app=leetcode id=621 lang=cpp
 *
 * [621] Task Scheduler
 */

// @lc code=start
class Solution {
public:

    int leastInterval(vector<char>& tasks, int n) {
      int total_tasks = tasks.size();
      if(!total_tasks) return 0;
      
      int intervals = 0;
      int arr_count[26] = {0};
      int max_freq = 0;
      int most_freq_count = 0;
      
      for(const auto &it : tasks) {
        arr_count[it-'A']++;
        if(max_freq == arr_count[it-'A']) {
          // this kind of char will have to insert between the max freq slot 
          most_freq_count++; 
        }
        else if(max_freq < arr_count[it-'A']) {
          max_freq = arr_count[it-'A'];
          most_freq_count = 1;
        }
      }

      intervals = (max_freq) * (n+1) - n + (most_freq_count-1);
      return max(intervals, total_tasks);
    }
};
// @lc code=end

