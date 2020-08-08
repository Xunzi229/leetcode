package main

import "fmt"

func main() {
	s := "ab"
	fmt.Println(s[1:2])

	fmt.Println(isPalindrome("aa"))
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}

	var start = 0
	var end = len(s) - 1
	for start <= end {
		if (len(s)%2 == 0) && (start+1 == end) && (s[start] == s[end]) {
			return true
		}

		if (len(s)%2 == 1) && (start == end) && (s[start] == s[end]) {
			return true
		}

		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return false
}
