/*
 * @lc app=leetcode id=767 lang=cpp
 *
 * [767] Reorganize String
 */

// @lc code=start
class Solution {
public:  
    string reorganizeString(string S) {
      if(S.empty()) return "";
      
      string ans = "";
      int arr_count[26] = {0};
      
      for(const auto &it : S) {
        arr_count[it-'a']++;
      }
      
      vector<pair<char, int>> vec_count;
      // unable to sort by map -> compile error
      // we turn to sort by vector
      for(int i=0; i<26; i++) {
        if(arr_count[i] > 0) {
          char insert_char = i+'a';
          vec_count.push_back(pair<char, int>(insert_char, arr_count[i]));
        }
      }
      
      /*
        1. sort the vector by value in ASC
        2. try to target the last char 
           (which has the larget frequency by the meantime)
        3. find separtor for the target char
      */
      while(vec_count.size()) {    
        sort(vec_count.begin(), vec_count.end(), compare);
        
        pair<char, int> t = vec_count[vec_count.size()-1];
        vec_count.pop_back();
        char t_char = t.first;
        int t_count = t.second;
        for(int i=0; i<t_count; i++) {
          ans+=t_char;
          if(i == t_count-1) 
            continue;
          // if not tail, try to get a separator
          char add_c = add_char(vec_count);
          // cannot get a char being separator
          if(add_c == 0 && i != t_count-1) 
            return "";
          ans+=add_c;
        }
      }
      
      return ans;
    }
  
private:
    char add_char(vector<pair<char, int>> &v) {
      if(v.size() == 0) return 0;
      char add_c = 0;
      v[v.size()-1].second --;
      add_c = v[v.size()-1].first;
      if(v[v.size() - 1].second == 0) v.pop_back();
      
      return add_c;
    }
      
    static bool compare(const pair<char,int> &a, const pair<int,int> &b) {
      return b.second > a.second;
    }
};
// @lc code=end

