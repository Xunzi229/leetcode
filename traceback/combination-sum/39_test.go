package combination_sum

import (
	"fmt"
	"testing"
)

/*
[2,3,6,7]
7

[2,3,5]
8

[2]
1

[1]
1

[1]
2

[2,7,6,3,5,1]
9
*/
func PackageCombinationSum() [][]int {
	return combinationSum([]int{2, 3, 5}, 8)
}

func TestPackageCombinationSum(t *testing.T) {
	fmt.Println(PackageCombinationSum())
}
