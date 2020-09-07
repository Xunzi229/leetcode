package maximum_subarray

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(a))
}
