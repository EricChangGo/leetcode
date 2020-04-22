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
      int buy_idx = 0;
      int sale_idx = 0;
      /**
        1. better min -> change holding stock, stock_idx = (i,i)
        2. new stock (old buy value <value < old sale value) -> sale the old one, stock_idx = (i,i)
        3. larger sale day -> recount the profit
      */
      for(int i=1; i<prices.size(); i++) {
        if(prices[i] > prices[sale_idx]) 
          t_profit += (prices[i] - prices[sale_idx]);
        else 
          buy_idx = i;
        sale_idx = i;
      }
      return t_profit;
    }
};
// @lc code=end

