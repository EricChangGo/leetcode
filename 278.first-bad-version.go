/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */

// @lc code=start
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

 func firstBadVersion(n int) int {
	/**
			find the turning point (false <-> true)
	*/
	start := 1
	end := n
	if isBadVersion(start) {
			// bad ...
			return start
	}
	
	badVersion := -1
	check := 0
	
	for true {
			check = (start+end)/2
			fmt.Println(check, start, end)
			if start == check {
					badVersion = end
					break
			}
			
			if isBadVersion(check) == true {
					// find first bad version in previous version
					end = check
			} else {
					// find first bad version in later version
					start = check
			}
	}
	return badVersion
}
// @lc code=end

