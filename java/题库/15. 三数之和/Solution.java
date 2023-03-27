/*
* 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/3sum
* */

class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);

        List<List<Integer>> ans = new ArrayList<List<Integer>>();

        int x = 0;
        while (x < nums.length) {
            if (nums[x] > 0) {
                break;
            }

            if (x > 0 && nums[x] == nums[x - 1]) {
                x++;
                continue;
            }

            int y = x + 1;
            if (y < nums.length && nums[x] + nums[y] > 0) {
                break;
            }

            int z = nums.length - 1;

            int target = -nums[x];

            for (; y < nums.length; ++y) {
                if (y > x + 1 && nums[y] == nums[y - 1]) {
                    continue;
                }
                while (y < z && nums[y] + nums[z] > target) {
                    --z;
                }
                if (y == z) {
                    break;
                }
                if (nums[y] + nums[z] == target) {
                    List<Integer> list = new ArrayList<Integer>();
                    list.add(nums[x]);
                    list.add(nums[y]);
                    list.add(nums[z]);
                    ans.add(list);
                }
            }
            ++x;
        }
        return ans;
    }
}