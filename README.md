# LeetCode solutions
Writing down solution ideas for some interesting questions.
Search the question ID as (-777-) to find the notes

# -53- maximum-subarray
There will be two method to solve this problem - Iterative and Divide&conquer

## Iterative
First we provide a iterative way to solve this problem.
```GO
func max(a,b int) int {
	if b > a {
		return b
	}
	return a
}

func maxSubArray(nums []int) int {
	MaxInt:= (1 << 31)-1
    MinInt:= -MaxInt-1
    
	temp:=0
    maxValue:=MinInt
    for i:=0; i<len(nums); i++ {
        if temp < 0 {
            temp = 0
        }
        temp+=nums[i]
        maxValue=max(maxValue, temp)
    }
    return maxValue
}
```
### ***Thinking process***
Here is a given input. Let's see how the code works.
```
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

- First, a naive thought is to add the numbers.
  - Pick the temparary total value (temp_sum) as sum, if it is better the choice

```Go
maxValue=max(maxValue, temp)
```

- Second, if a temparary total value is negative
  - It will be a great idea to forgive the temp sub string. In other words, this sub string is not a good choice, because it will turn down your value.

### ***Key steps***
1. sub-string = [-2, 1, -3], temp = 1+(-3) → temp = 0 **(drop)**
2. sub-string = [4] **(good choice)**  → maxValue = 4
3. sub-string = [4, -1, __, __ ...], maxValue = 4
4. In step 3, depends on the further sub __, __ ...
    - If the sub will turn down then the temp will never replace the maxValue.
    - Or else, might get larger temp in order to replace maxValue.

## Divide&conquer
```Go
const MaxInt = (1 << 31) - 1
const MinInt = -MaxInt - 1

func max(a,b int) int {
	if b > a {
		return b
	}
	return a
}

func getLeftMax(num []int, start int, end int) int {
	temp:=0
	maxValue:=MinInt
	for i:=end; i>=start; i-- {
		temp+=num[i]
		maxValue=max(maxValue, temp)
	}
	return maxValue
}

func getRightMax(num []int, start int, end int) int {
	temp:=0
	maxValue:=MinInt
	for i:=start; i<=end; i++ {
		temp+=num[i]
		maxValue=max(maxValue, temp)
	}
	return maxValue
}

func maxSubArray(nums []int) int {
	return getMaxSubs(nums, 0, len(nums)-1)
}

func getMaxSubs(nums []int, start int, end int) int {
	if start == end {
		return nums[start]
	}
	mid := (start+end)/2
	leftSub := getMaxSubs(nums, start,  mid)
	rightSub := getMaxSubs(nums, mid+1, end)

	leftMax := getLeftMax(nums, start, mid)
	rightMax := getRightMax(nums, mid+1, end)
	total := leftMax + rightMax
	return max(max(leftSub, rightSub), total)
}
```

### ***Thinking process***

The idea is to divide the array into two half (left, right) and try to pick the maxValue in each division method recursively.

The main flow is to start at the middle for the initial sub string, then decide to extended or repicked a better result in right or left.

Given the same input down below.
```
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

1. First picked the -1 in the middle, and seperate the array to [-2,1,-3,4,-1] and [2,1,-5,4].
2. Target the former one [-2,1,-3,4, 1], then we try to focus on a smaller divsions - left [-2,1] and right [4,-1]

3. Right maxValue will be $max(4, 4+(-1)) = 4$, how about the left value?

***Most tricky part*** - decide the value of the right.
4. Left value will be $max(1, 1+(-2))$. The direction adding each value should be right to left. In this direction, we are able to extend the middle substring.
    - If we go left to right, $max(-2, 1+(-2))$ will be a total incorrect value for extending middle substring.
    - If we picked the wrong way, we will miss the value by merging middle value with left value.

5. Merge or keep searching? 
    - Merge: $total = left maxValue + right maxValue$  
    - Keep searching by recursive left right sub string

6. Choose the max result $return max(max(leftSub, rightSub), total)$

# -123- Best Time to Buy and Sell Stock III 
## Basic Info
Similar with Problem 122, but **You may complete at most two transactions.***


```
Input: [3,3,5,0,0,3,1,4]
Output: 6
Explanation: 
    Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
    Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
```

## Solution
```C++
class Solution {
public:
    int maxProfit(vector<int>& prices) {
      int trans_limit = 2;
      vector<pair<int, int>> records = {pair<int, int>(INT_MIN,0), pair<int, int>(INT_MIN,0)};;
          
      for(int i=0; i<prices.size(); i++) {
        records[0].first = max(-prices[i], records[0].first);
        records[0].second = max(records[0].first + prices[i] , records[0].second);
        
        for(int t=1; t<trans_limit; t++) {
          records[t].first = max(records[t-1].second - prices[i], records[t].first);
          records[t].second = max(records[t].first + prices[i], records[t].second);
        }
      }
      return records[trans_limit-1].second;
    }
};
```

## Thinking Process
- The idea is to storing 2 most highest profit by dp.
So, every time we picked a new target i, we get the new result by taking comparison with the previous result.

- The key part is to control this two records properly. 
### First record
- records[0].first represents our latest buying transaction
  1. consider buying the target stock $i → -price[i]$
  2. $pass$

- records[1].second represents our latest saling transaction
  1. consider saling the target stock $i$ to be our first transaction $record → records[t].first + prices[i]$
  2. $pass$
- Picked the best decision, smaller for buy, larger for sale → using max function for both action
### Second record
- same concept above except need to consider the profit from latest transaction
- second buy record → $records[t-1].second - prices[i]$ $or$ $pass$
- if there is no second transaction, this record value will be copy from the first record.

## Example
```
Input: [3,3,5,0,0,3,1,4]
Output: 6
```
### Step 1

$record[0] = [3,\widehat{3,5},0,0,3,1,4]$

$record[1] = [3,\widehat{3,5},0,0,3,1,4]$

$profit = 2

### Step 2

record[0] = $[3,3,5,0,\widehat{0,3},1,4]$

record[1] = $[3,\widehat{3,5},0,0,3,1,4]$

$profit = 5

### Step 3

record[0] = $[3,3,5,0,\widehat{0,3,1,4}]$

record[1] = $[3,\widehat{3,5},0,0,3,1,4]$

$profit = 6

# -621- Task Scheduler 
## Basic Info
There is a non-negative cooling interval n that means between two same tasks, there must be at least n intervals that CPU are doing different tasks or just be idle.

You need to return the least number of intervals the CPU will take to finish all the given tasks.

![Example](https://i.imgur.com/18P3qUi.png)

* Approximate with 767 
* We get 50% faster only by using the similar algo with 767 - see github 621_1.0
* Try to figure a new algo which only focus on the max freqency char, since idle could be insert,.

## Solution
```C++
class Solution {
public:

    int leastInterval(vector<char>& tasks, int n) {
      int total_tasks = tasks.size();
      if(!total_tasks) return 0;
      
      int intervals = 0;
      int arr_count[26] = {0};
      int max_freq = 0;
      int most_freq_count = 1;
      
      for(const auto &it : tasks) {
        arr_count[it-'A']++;
        if(max_freq == arr_count[it-'A']) {
          // this kind of char will have to insert between the max freq slot 
          most_freq_count++; 
        }
        if(max_freq < arr_count[it-'A']) {
          max_freq = arr_count[it-'A'];
          most_freq_count = 1;
        }
      }
      
      intervals = (max_freq) * (n+1) - n + (most_freq_count-1);
      
      return intervals > total_tasks ? intervals : total_tasks;
    }
};
```
## Thinking Process

### Step1
```
tasks = [A,A,A,B,B,B], n=2
max_freq = 3 
max_freq_set = {A,B} 
max_freq_count = 2
```
Let's start design an simple algo.

1. Build the max freq string 
     
     - $ (maxfreq)*(n+1) → [AiiAiiAii] $ 
  
2. Cut the redundant tail idle -> at most n idle need cut 
     
      - $ -n → [AiiAiiA] $
3. We need to insert other most freq char into tail and previous max freq char(A) 
 
      - $ (mostFreqCount-1) → [ABiABiAB] $

*** Q. Any blind side above? ***
  
  Without considering other characters thats not belong to ***max_freq_set***.

### Step 2 - Consider other characters
```
tasks = [A,A,A,B,B,C,D], n=2
other_chars = {B,C,D}
other_chars_freq = 4 
```

  In this case, we consider more about inserting other    characters. Since max_freq_count = 1, left space is all for max_freq_chars.
        
  - $[AiiAiiAii] → [ABiABiACi] → [ABDiABiACi]$
        
As a result, the upper algo seems just fine. 
> 
*** Q. But what if the space is not enough to fill in? ***

### Step 3 - not enough space?
```
tasks = [A,A,A,B,B,C,D,E,F,F], n=2
max_freq = 2
other_chars = {B,C,D,E,F}
other_chars_freq = 6 
```

1. Same step above
   - $[AiiAiiAii] → [ABiABiACi] → [ABCABDACE]$
   - ***What about $F$?*** 
2. Stick it next to max_chars
   - According to the rules, F need to be seperate with 2 blocks. So, let just stick it next to A. Accordingly, this would not cause any idles.

3. Output ***Size of Tasks***
   

### Step 4 - Find the condition
There will be only two situation, enough space or not.

> Q. How do we output the correct answer by condition?

#### ***This is the toughest part in this problem***

* Enough space -  ***Step 1 Output > Size of Tasks (Step 3 Output)*** 
* Not enough space - ***Size of Tasks (Step 3 Output) > Step 1 Output ***

> Enough space
>   - idles will be set between the chars

> Not enough space
>   - after fill in all the spaces, the string grows.


### Output Solution - max(step1, Size of Tasks)

### More Questions
> Q. What about mutiple max chars, max_freq_count = m
>    - If n <= m, using step1, spaces is enough for  max_freq_set elements all sticking together.
>    - Else, not enough space. Go for step 3.

# -912- sort-an-array

#### ***Merge Sort***
```C++
class Solution {
// practice merge sort
public:
    vector<int> sortArray(vector<int>& nums) {
      merge_sort(nums, 0, nums.size()-1);
      return nums;
    }
    
    void merge_sort(vector<int>& nums, int start, int end) {
      // stop when the we target the basic one element
      if(start >= end) return; 
      
      // divide the array into half
      int mid = (start+end)/2;
      // sort left elements
      merge_sort(nums, start, mid);
      // sort right elements
      merge_sort(nums, mid+1, end);
      
      merge(nums, start, mid, end);
    }
    
    void merge(vector<int>& nums, int start, int mid, int end) {
      /*
        1. Set left <= mid, right > mid
        2. Arrange into a new array with all left right element
        
        Ex: left = [1,5], right = [3,8]
        New array = [1,3,5,8]
      */
      vector<int> left(nums.begin()+start, nums.begin()+mid+1);
      vector<int> right(nums.begin()+mid+1, nums.begin()+end+1);
      left.push_back(INT_MAX);
      right.push_back(INT_MAX);
      
      int l_idx = 0;
      int r_idx = 0;
      
      for(int i=start; i<=end; i++) {
        // set nums[i] 
        if(left[l_idx] <= right[r_idx]) {
          nums[i] = left[l_idx];
          l_idx++;
        }
        else {
          nums[i] = right[r_idx];
          r_idx++;
        }
      }
    }
  
    void print_list(vector<int>& list) {
      for(const auto &it: list) {
        cout << it << " ";
      }
      cout << endl;
    }
};
```



