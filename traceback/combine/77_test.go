package combine

import (
	"fmt"
	"testing"
)

func PackageCombine() [][]int {
	return combine(4, 2)
}

func TestPackageCombine(t *testing.T) {
	fmt.Println(PackageCombine())
}
