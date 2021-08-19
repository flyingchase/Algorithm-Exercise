package HashSet;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

/*给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。


 */
public class E_1002_commonChars {
    public static void main(String[] args) {
        String[] words = {"bella", "label", "roller"};
        System.out.println(new E_1002_commonChars().commonChars(words).toString());
    }

    public List<String> commonChars(String[] words) {
        List<String> res = new ArrayList<String>();
        if (words == null || words.length < 1) {
            return res;
        }
        HashMap<Character, Integer> base = new HashMap<>();

        for (int i = 0; i < words[0].length(); i++) {
            base.put(words[0].charAt(i), base.getOrDefault(words[0].charAt(i), 0) + 1);
        }

        for (int i = 1; i < words.length; i++) {
            String word = words[i];
            char[] chars = word.toCharArray();
            HashMap<Character, Integer> map = new HashMap<>();
            for (char c : chars) {
                if (base.containsKey(c)) {
                    map.put(c, Math.min(base.get(c), map.getOrDefault(c, 0) + 1));
                }
            }
            base = map;
        }

        for (Character character : base.keySet()) {
            for (Integer i = 0; i < base.get(character); i++) {
                res.add(character +"");
            }
        }
        return res;
    }

}
