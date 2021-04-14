package unique_paths_iii

import (
	"fmt"
	"testing"
)

func PackageSubSetsWithDup(dy [][]int) int {
	return uniquePathsIII(dy)
}

//2 2 1 [
//[-1 -1 -1 0]
//[-1 -1 -1 0]
//[-1 -1 -1 -1]
//]

func TestPackageSubSetsWithDup(t *testing.T) {
	dy := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1}}
	fmt.Println(PackageSubSetsWithDup(dy))
}
