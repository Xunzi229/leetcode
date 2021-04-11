package subsets_ii

import (
	"fmt"
	"testing"
)

func PackageSubSetsWithDup(dy []int) [][]int {
	return subsetsWithDup(dy)
}

func TestPackageSubSetsWithDup(t *testing.T) {

	fmt.Println(PackageSubSetsWithDup([]int{1, 2, 2}))
}
