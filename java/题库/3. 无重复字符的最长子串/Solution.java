/*
3. 无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
https://leetcode.cn/problems/longest-substring-without-repeating-characters/
* * */

class Solution {
    public int lengthOfLongestSubstring(String s) {
        int max = 0;
        List<Byte> window = new ArrayList<>();
        for (char c : s.toCharArray()) {
            for (int i = 0; i < window.size(); i++) {
                if (window.get(i) == c) {
                    window = window.subList(i + 1, window.size());
                    break;
                }
            }
            if (max <= window.size()) {
                max = window.size() + 1;
            }
            window.add((byte) c);
        }
        return max;
    }
}