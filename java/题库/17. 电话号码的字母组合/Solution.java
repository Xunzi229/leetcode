/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
* * */


import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Solution {
    public List<String> letterCombinations(String digits) {
        Map<Character, ArrayList<String>> m = new HashMap<Character, ArrayList<String>>() {
            {
                put('2', new ArrayList<String>() {{
                    add("a");
                    add("b");
                    add("c");
                }});
            }

            {
                put('3', new ArrayList<String>() {{
                    add("d");
                    add("e");
                    add("f");
                }});
            }

            {
                put('4', new ArrayList<String>() {{
                    add("g");
                    add("h");
                    add("i");
                }});
            }

            {
                put('5', new ArrayList<String>() {{
                    add("j");
                    add("k");
                    add("l");
                }});
            }

            {
                put('6', new ArrayList<String>() {{
                    add("m");
                    add("n");
                    add("o");
                }});
            }

            {
                put('7', new ArrayList<String>() {{
                    add("p");
                    add("q");
                    add("r");
                    add("s");
                }});
            }

            {
                put('8', new ArrayList<String>() {{
                    add("t");
                    add("u");
                    add("v");
                }});
            }

            {
                put('9', new ArrayList<String>() {{
                    add("w");
                    add("x");
                    add("y");
                    add("z");
                }});
            }
        };

        char[] bytes = digits.toCharArray();

        List<Character> result = new ArrayList<>();

        for (Character num : bytes) {
            if (num != '1') {
                result.add(num);
            }
        }
        List<String> list = new ArrayList<>();
        int index = 0;
        while (index < result.size()) {
            List<String> l = new ArrayList<>();

            for (int y = 0; y < m.get(result.get(index)).size(); y++) {
                if (index > 0) {
                    for (String s : list) {
                        l.add(s + m.get(result.get(index)).get(y));
                    }
                } else {
                    l = m.get(result.get(index));
                }
            }
            list = l;
            index++;
        }
        return list;
    }
}