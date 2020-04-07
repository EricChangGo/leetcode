--
-- @lc app=leetcode id=176 lang=mysql
--
-- [176] Second Highest Salary
--

-- @lc code=start
-- Write your MySQL query statement below
-- 1. Select distinct Salary Order by descending, and use LIMIT to get the second one
--    - prefer not to use ranking in this situation (Q: Will it be faster?)
--    - distinct won't able to fit for the problem extention, etc: thrird greatest number
-- 2. Check if the row result is empty -> yes, return null
SELECT IFNULL( (SELECT DISTINCT Salary FROM Employee ORDER BY Salary DESC LIMIT 1,1), NULL) as SecondHighestSalary;
-- @lc code=end

