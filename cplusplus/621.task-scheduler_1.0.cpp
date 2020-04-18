/*
 * @lc app=leetcode id=621 lang=cpp
 *
 * [621] Task Scheduler
 */

// @lc code=start
class Solution {
public:
    /* 
      approximate with 767 
      no need to use vector, no need to output string
    */
    int leastInterval(vector<char>& tasks, int n) {
      if(!tasks.size()) return 0;
      
      string schedule_str = "";
      int task_count = 0;
      int arr_size = 26;
      int arr_count[26] = {0};
      
      for(const auto &it : tasks) {
        arr_count[it-'A']++;
      }
      
      // sort in DESC
      sort(arr_count, arr_count+arr_size, greater<int>());
      int target_index = 0;
      
      while(target_index < arr_size) {
        // char t_char=t.first;
        int t_count = arr_count[target_index];
        target_index ++;
        
        for(int i=0; i<t_count; i++) {
          // schedule_str+=t_char;
          task_count ++;
          if(i == t_count-1) 
            break;
          
          int insert_candidate = target_index;
          for(int j=0; j < n; j++) {
            // if not tail, try to get a char to insert, or idle
            if(target_index < arr_size  && insert_candidate < arr_size) {
              arr_count[insert_candidate] --;
              // schedule_str+=vec_count[insert_candidate].first;
            }
            task_count++;
            insert_candidate++;
          }
          sort(arr_count+target_index, arr_count+arr_size, greater<int>());
        }
      }
      return task_count;
    }
  
    static bool compare(const pair<char,int> &a, const pair<int,int> &b) {
      return b.second < a.second;
    }
};
// @lc code=end

