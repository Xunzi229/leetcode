import java.util.*;

/*
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，
并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/two-sum
 */

class Solution {
    public int[] twoSum1(int[] nums, int target) {
        Map<Integer, Integer> m = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if (m.containsKey(target - nums[i])) {
                return new int[]{m.get(target - nums[i]), i};
            }
            m.put(nums[i], i);
        }
        return new int[]{};
    }

    public int[] twoSum1(int[] nums, int target) {
        for (int i = 0; i < nums.length; i++) {
            if (i + 1 > nums.length) {
                return new int[]{};
            }
            for (int y = i + 1; y < nums.length; y++) {
                if (nums[i] + nums[y] == target && nums[i] != nums[y]) {
                    return new int[]{i, y};
                }
            }
        }
        return new int[]{};
    }
}