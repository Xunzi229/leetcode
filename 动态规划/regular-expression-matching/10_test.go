package regular_expression_matching

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	//fmt.Println(isMatch("abb", "ab*")) except true
	//fmt.Println(isMatch("abb", "ab*"))
	//fmt.Println(isMatch("aa", "a*"))
	//fmt.Println(isMatch("ab", ".*"))
	//fmt.Println(isMatch("aab", "c*a*b"))

	if isMatch("mississippi", "mis*is*p*.") == false {
		fmt.Println("pass")
	}
}
