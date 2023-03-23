// 给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。
// 找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/container-with-most-water

class Solution {
    public int maxArea(int[] height) {
        int max = 0;

        for (int i = 0; i < height.length; i++) {
            for (int y = i + 1; y < height.length; y++) {
                if (height[y] > height[i]) {
                    if (height[i] * (y - i) > max) {
                        max = height[i] * (y - i);
                    }
                } else {
                    if (height[y] * (y - i) > max) {
                        max = height[y] * (y - i);
                    }
                }
            }
        }
        return max;
    }

    public int maxArea1(int[] height) {
        int max = 0;

        int left = 0;
        int right = height.length - 1;

        while (left != right) {
            if (height[left] < height[right]) {
                if ((right - left) * height[left] > max) {
                    max = (right - left) * height[left];
                }
                left++;
                continue;
            }
            if (height[left] >= height[right]) {
                if ((right - left) * height[right] > max) {
                    max = (right - left) * height[right];
                }
                right--;
            }
        }
        return max;
    }
}