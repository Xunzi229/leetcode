//https://leetcode-cn.com/problems/first-bad-version/
package first_bad_version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	start, end := 1, n

	for start+1 < end {
		mid := start + (end-start)/2
		if isBadVersion(mid) {
			for {
				if !isBadVersion(mid - 1) {
					return mid
				}
				mid--
			}
		}
		start = mid + 1
	}

	if isBadVersion(start) {
		return start
	}
	if isBadVersion(end) {
		return end
	}

	return 0
}

func isBadVersion(version int) bool { return false }
