package String;

import java.lang.invoke.VarHandle;
import java.time.temporal.Temporal;
import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.HashSet;

/*给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。

元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。*/
public class E_345_reverseVowels {

    public static void main(String[] args) {
        String s = "leetcode";
        System.out.println(new E_345_reverseVowels().reverseVowels(s));
    }

    public String reverseVowels(String s) {
        if(s == null || s.length()==0) return s;
        String vowels = "aeiouAEIOU";
        // 将字符串转化成char类型数组
        char[] chars = s.toCharArray();
        int start = 0;
        int end = s.length()-1;
        while(start < end){
            // 双指针相向而行找元音字符
            while(start < end && !vowels.contains(chars[start]+"")){
                start++;
            }
            while(start < end && !vowels.contains(chars[end]+"")){
                end--;
            }
            char temp = chars[start];
            chars[start] = chars[end];
            chars[end] = temp;
            start++;
            end--;
        }
        return new String(chars);
    }
}
