package numbers_with_same_consecutive_differences

import (
	"fmt"
	"testing"
)

func PackageNumsSameConsecDiff(n, k int) []int {
	return numsSameConsecDiff(n, k)
}

func TestPackageNumsSameConsecDiff(t *testing.T) {
	fmt.Println(PackageNumsSameConsecDiff(3, 7))
	fmt.Println(PackageNumsSameConsecDiff(2, 1))
	fmt.Println(PackageNumsSameConsecDiff(2, 0))
	fmt.Println(PackageNumsSameConsecDiff(2, 2))
}
