package beautiful_arrangement

import (
	"fmt"
	"testing"
)

func PackageCountArrangement(index int) int {
	return countArrangement(index)
}

func TestPackageCountArrangement(t *testing.T) {
	fmt.Println(PackageCountArrangement(1)) // 1
	fmt.Println(PackageCountArrangement(2)) // 2
	fmt.Println(PackageCountArrangement(3)) // 3
	fmt.Println(PackageCountArrangement(4)) // 8
	fmt.Println(PackageCountArrangement(5)) // 10
	fmt.Println(PackageCountArrangement(6)) // 36
	fmt.Println(PackageCountArrangement(7)) // 41
}
