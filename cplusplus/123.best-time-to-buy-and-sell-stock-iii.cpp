/*
 * @lc app=leetcode id=123 lang=cpp
 *
 * [123] Best Time to Buy and Sell Stock III
 */

// @lc code=start
class Solution {
public:
    int maxProfit(vector<int>& prices) {
      int trans_limit = 2;
      vector<pair<int, int>> records = {pair<int, int>(INT_MIN,0), pair<int, int>(INT_MIN,0)};;
          
      for(int i=0; i<prices.size(); i++) {
        // first buy -> buy the min cost, not buy -> continue empty
        records[0].first = max(-prices[i], records[0].first);
        // sell -> sell in the max profit, not sell -> continue holding
        records[0].second = max(records[0].first + prices[i] , records[0].second);
        
        for(int t=1; t<trans_limit; t++) {
          records[t].first = max(records[t-1].second - prices[i], records[t].first);
          records[t].second = max(records[t].first + prices[i], records[t].second);
        }
      }
      return records[trans_limit-1].second;
    }
};
// @lc code=end

