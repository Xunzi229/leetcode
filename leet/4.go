//https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

package main

import (
    "context"
    "fmt"
    "time"
)

//func main(){
//    nums1 := []int{1, 2}
//    nums2 := []int{3, 4}
//
//    fmt.Println(findMedianSortedArrays(nums1, nums2))
//}
//
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//    context.Context()
//    return 0
//}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	go handle(ctx, 1500 * time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
  default:
    time.Sleep(time.Millisecond * 1500)

		fmt.Println("process request with", duration)
	}
}

/*
// 自己的辣鸡解法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var nums3 []int
    nums3 = append(nums1, nums2...)
    sort.Ints(nums3)

    Len := len(nums3)
    if Len % 2 == 0{
        right := Len / 2
        left := right - 1
        return (float64(nums3[left]) + float64(nums3[right]) ) / 2.0
    } else {
        left := (Len - 1) / 2
        return float64(nums3[left])
    }
}
*/
