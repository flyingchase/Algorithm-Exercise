package HashSet;

import java.util.HashMap;

@SuppressWarnings({"all"})
/*给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；否则，返回 false。*/
public class E_953_isAlientSorted {
    HashMap<Character, Integer> map;

    public boolean isAlienSorted(String[] words, String order) {
        map = new HashMap<Character, Integer>();
        char[] ch = order.toCharArray();
        for (int i = 0; i < ch.length; i++) {
            map.put(ch[i], i++);
        }

        for (int i = 0; i < words.length-1; i++) {
            if (!compare(words[i],words[i+1])) {
                return false;
            }
        }
        return true;
    }

    private boolean compare(String s1, String s2) {
        int l1 = s1.length();
        int l2 = s2.length();
        for (int i = 0,j = 0; i < l1&&j<l2; i++,j++) {
            if (s1.charAt(i)!=s2.charAt(j)) {
                if (map.get(s1.charAt(i))>map.get(s2.charAt(j))) {
                    return false;
                }else {
                    return true;
                }
            }
        }
        if (l1>l2) {
            return false;
        }
        return true;
    }
}
