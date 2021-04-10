package permutations

import (
	"fmt"
	"testing"
)

func PackagePermute() [][]int {
	return permute([]int{1, 2, 3})
}

func TestPackagePermute(t *testing.T) {
	fmt.Println(PackagePermute())
}
