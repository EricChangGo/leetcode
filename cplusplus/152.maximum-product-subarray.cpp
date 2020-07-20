  // Thinking process
  // 1. Simple interation 
  //     a. temp max  
  //     b. save as result
  //     c. continue to find temp max
	// 	```
	// 		tempMax = max({oldValue*nums[i], nums[i]});
	// 		ans = max(tempMax, ans);
	// 		// cannot solve [-2, 3, -1]
	// 	```
  // 2. Seems like we have to consider the previous temp max, which the numbers are consistent with num[i]
  // 3. Since our num[i] might be negative, we get prev temp max might be positive or negative 
	//   > Record both temporary minimum and maximum
	// 	```
	// 		tempMin = min({oldValue*nums[i], nums[i], tempMin*nums[i]});
	// 	```
class Solution {
public:
    int maxProduct(vector<int>& nums) {
      if(!nums.size()) return 0;
      int ans = nums[0];
      int tempMax = nums[0];
      int tempMin = nums[0];
      int oldValue;
      
      for(int i=1; i < nums.size(); i++) {
        oldValue = tempMax;
        tempMax = max({oldValue*nums[i], nums[i], tempMin*nums[i]});
        tempMin = min({oldValue*nums[i], nums[i], tempMin*nums[i]});
        ans = max(tempMax, ans);
      }
      return ans;
    }
};