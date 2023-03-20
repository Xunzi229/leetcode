/*
 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/group-anagrams

这个问题解决的就是相同字符的字符串（顺序不同）分组后在同一个数据中
*/

class Solution {
  public List<List<String>> groupAnagrams(String[] strs) {

    // 具有相同字符的在同一个数组中
    Map<String, List<String>> m = new HashMap<String, List<String>>();

    for (String str : strs) {
      char[] chs = str.toCharArray();

      Arrays.sort(chs);
      String key = new String(chs);
      List<String> list = m.getOrDefault(key, new ArrayList<String>());
      list.add(str);
      m.put(key, list);
    }
    return new ArrayList<List<String>>(m.values());
  }
}