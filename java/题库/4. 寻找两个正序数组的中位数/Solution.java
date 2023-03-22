package So;
/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/median-of-two-sorted-arrays
*/

class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int nX = 0;
        int nY = 0;

        int total = nums1.length + nums2.length;
        if (total == 0) {
            return 0;
        }

        if (total == 1) {
            if (nums1.length == 1) {
                return nums1[0];
            }
            return nums2[0];
        }

        List<Integer> list = new ArrayList<>();

        for (int i = 0; i < total; i++ ) {
            if (nX >= nums1.length) {
                list.add(nums2[nY]);
                nY++;
            } else if (nY >= nums2.length) {
                list.add(nums1[nX]);
                nX++;
            } else if (nums1[nX] < nums2[nY]) {
                list.add(nums1[nX]);
                nX++;
            } else {
                list.add(nums2[nY]);
                nY++;
            }
        }

        if (total % 2 == 1) {
            return list.get(total / 2);
        } else  {
            return (list.get((total / 2) - 1) + list.get((total / 2)))  / 2.0;
        }
    }


    public static double findMedianSortedArrays2(int[] nums1, int[] nums2) {
        int nX = 0;
        int nY = 0;

        int total = nums1.length + nums2.length;
        if (total == 0) {
            return 0;
        }

        if (total == 1) {
            if (nums1.length == 1) {
                return nums1[0];
            }
            return nums2[0];
        }

        int upNum = 0;
        int centerIndex = total / 2;

        int c = 0;
        for (int i = 0; i < total; i++ ) {
            if (nX >= nums1.length) {
                c = nums2[nY];
                nY++;
            } else if (nY >= nums2.length) {
                c = nums1[nX];
                nX++;
            } else if (nums1[nX] < nums2[nY]) {
                c = nums1[nX];
                nX++;
            }else {
                c = nums2[nY];
                nY++;
            }
            if (centerIndex == i) {
                if (total % 2 == 1) {
                    return c;
                } else  {
                    return (upNum + c)  / 2.0;
                }
            }
            upNum = c;
        }
        return 0;
    }
}