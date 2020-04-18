# LeetCode solutions
Writing down solution ideas for some interesting questions.
Search the question ID as (-777-) to find the notes

# -621-
## Basic Info
There is a non-negative cooling interval n that means between two same tasks, there must be at least n intervals that CPU are doing different tasks or just be idle.

You need to return the least number of intervals the CPU will take to finish all the given tasks.

![Example](https://i.imgur.com/18P3qUi.png)

* Approximate with 767 
* We get 50% faster only by using the similar algo with 767 - see github 621_1.0
* Try to figure a new algo which only focus on the max freqency char, since idle could be insert,.

## Solution
```
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
     
     - $(maxfreq)*(n+1) → [AiiAiiAii]$ 
  
2. Cut the redundant tail idle -> at most n idle need cut 
     
      - $-n → [AiiAiiA]$
3. We need to insert other most freq char into tail and previous max freq char(A) 
 
      - $(mostFreqCount-1) → [ABiABiAB]$

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
other_chars = {B,C,D,E,F}, other_chars_freq = 6 
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
>   - after filled in all the spaces, the string grows.


### Output Solution - max(step1, Size of Tasks)

### More Questions
> Q. What about mutiple max chars, max_freq_count = m
>    - If n <= m, using step1, spaces is enough for  max_freq_set elements all sticking together.
>    - Else, not enough space. Go for step 3.



