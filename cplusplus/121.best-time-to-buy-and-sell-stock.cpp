/*
 * @lc app=leetcode id=121 lang=cpp
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
class Solution {
public:
    int maxProfit(vector<int>& prices) {
      if(!prices.size()) return 0;
      
      int profit=0;
      // record the lowest price day
      int min_value=prices[0];
      
      for(int t=1; t<prices.size(); t++) {
        // update profit
        if(prices[t] - min_value > profit)
          profit = prices[t] - min_value;
        
        // check better min value for all element
        if(min_value > prices[t])
          min_value = prices[t];
      }
        
      return profit;
    }
};
// @lc code=end

