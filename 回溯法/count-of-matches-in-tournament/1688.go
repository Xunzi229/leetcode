/*
https://leetcode-cn.com/problems/count-of-matches-in-tournament/
*/
package count_of_matches_in_tournament

func numberOfMatches(n int) int {
	if n <= 1 {
		return 0
	}
	if n%2 == 0 {
		t2 := numberOfMatches(n / 2)
		return n/2 + t2
	}
	t1 := numberOfMatches((n-1)/2 + 1)
	return (n-1)/2 + t1
}
