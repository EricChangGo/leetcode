func maxProfit(prices []int, fee int) int {
	/** 
			Problem III method need to be fixed, 
			since we have no idea how large can the number of transaction will be
	*/
	/** 
			Go back to use greedy of Problem II or use Problem III by changing it to DP
	*/
	
	/** 
			Greedy
			1. larger sale day (contain) -> recount the profit
			2. Key Part - when i value < sale value (there is an old sale point)
					a. new stock (i value < sale value) -> sale the old one, set new (buy,sale) = (i,i)
					b. profit has been consume by fee (i value < buy value ) -> drop the old holding, set new (buy,sale) = (i,i)
					Waring: b must be after a, or else we will immediately drop our holding without considering profit
	*/
	
	total_profit := 0;
	sale_idx := 0;
	buy_idx := 0;
	for i:=1; i<len(prices); i++ {
			if prices[i] >= prices[sale_idx] {
					// change better sale date
					sale_idx = i
			} else {
					/** 
							prices[i] < prices[sale_idx]
							
							Key Part - should we sale it or not?
							sale: temp_profit
							contain: prices[i] - buy = possible_profit
							
							ex.[2,4,1,5]
					*/
					temp_profit := prices[sale_idx] - prices[buy_idx] - fee
					possible_profit := prices[i] - prices[buy_idx]
					if temp_profit > 0 && temp_profit > possible_profit {
							// sale holding and buy new stock i
							total_profit += temp_profit
							buy_idx = i
							sale_idx = i
					} else if prices[i] < prices[buy_idx] {
							// no sale point, just buy new stock i
							buy_idx = i
							sale_idx = i
					}
					// keep holding
			}
	}
	
	// add the last decision
	temp_profit := prices[sale_idx] - prices[buy_idx] - fee
	if prices[sale_idx] - prices[buy_idx] - fee > 0 {
			total_profit += temp_profit
	}
	return total_profit;
}