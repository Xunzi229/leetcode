package combination_sum_iii

import (
	"fmt"
	"testing"
)

func PackageCombinationSum3() [][]int {
	return combinationSum3(3, 9)
}

func TestPackageCombinationSum3(t *testing.T) {
	fmt.Println(PackageCombinationSum3())
}
