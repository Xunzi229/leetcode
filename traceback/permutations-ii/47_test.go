package permutations_ii

import (
	"fmt"
	"testing"
)

func PackagePermuteUnique() [][]int {
	return permuteUnique([]int{1, 1, 2})
}

func TestPackagePermuteUnique(t *testing.T) {
	fmt.Println(PackagePermuteUnique())
}
