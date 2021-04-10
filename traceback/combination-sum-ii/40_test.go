package combination_sum_ii

import (
	"fmt"
	"testing"
)

/*
[10,1,2,7,6,1,5]
8

[2,5,2,1,2]
5

*/

/*
[3,1,3,5,1,1]
8
*/
func PackageCombinationSum() [][]int {
	return combinationSum2([]int{3, 1, 3, 5, 1, 1}, 8)
}

func TestPackageCombinationSum(t *testing.T) {
	fmt.Println(PackageCombinationSum())
}
