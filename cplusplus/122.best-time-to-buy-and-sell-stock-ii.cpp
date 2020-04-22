/*
 * @lc app=leetcode id=122 lang=cpp
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
class Solution {
public:
    int maxProfit(vector<int>& prices) {  
      int t_profit = 0;
      int hold_stock = 0;
      vector<pair<int, int>> profits_idx;
      profits_idx.push_back(pair<int, int> (0, 0));
      
      for(int i=1; i<prices.size(); i++) {
        // update profit
        int target_value = prices[i];
        int max_add_value = 0;
        int exchange_idx = -1;

        /**
          Key: you must sell the stock before you buy again
          Try to arrange to sell this i day.
        */
        for(int j=0; j<profits_idx.size(); j++) {
          
          // check overlap
          if(profits_idx[hold_stock].second > profits_idx[j].second) {
            exchange_idx == -2;
            continue;
          }

          if(target_value > prices[profits_idx[j].first]) {
            int add_value = target_value - prices[profits_idx[j].second];
            if(add_value > max_add_value) {
              max_add_value = add_value;
              exchange_idx = j;
            }
          }
        }
        
        if(exchange_idx == -2) {
          profits_idx.clear();
          continue;
        }

        if(exchange_idx == -1) {
          profits_idx.clear();
          profits_idx.push_back(pair<int, int> (i, i));
        }
        else {
          // sale in day i, which buy from day exchange_idx
          hold_stock = exchange_idx;
          t_profit += (prices[i] - prices[profits_idx[exchange_idx].second]);
          profits_idx[exchange_idx].second = i;
        }
      }
      
      return t_profit;
    }
};
// @lc code=end

