package HashMap;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

/*给一字符串数组, 将 错位词(指相同字符不同排列的字符串) 分组


 错位分组
 */
public class M_49_groupAnagrams {
    public List<List<String>> groupAnagrams(String[] strs) {
        if (strs==null) {
            return null;
        }
        HashMap<String, List<String>> map = new HashMap<>();
        for (String str : strs) {
            char[] ch = str.toCharArray();
            Arrays.sort(ch);
            String s = String.valueOf(ch);
            map.putIfAbsent(s,new ArrayList<>());
            map.get(s).add(s);
        }
        return  new ArrayList<>(map.values());
    }
}
