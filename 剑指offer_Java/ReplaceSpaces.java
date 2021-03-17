package 剑指offer_Java;
import 剑指offer_Java.*;


public class ReplaceSpaces {
    //方法一: 遍历替换
    public String replaceSpace(StringBuffer str) {
        if (str == null || str.length() == 0) {
            return str.toString();
        }
        StringBuilder sb = new StringBuilder();
        int len = str.length();
        for (int i = 0; i < len; ++i) {
            char ch = str.charAt(i);
            sb.append(ch == ' ' ? "%20" : ch);
        }

        return sb.toString();
    }

    //方法二: 双指针

    public String replaceSpace1(StringBuffer str){
        if(str == null || str.length()==0)  return str.toString();
        int len = str.length();
        for(int i = 0;i<len;i++){
            if(str.charAt(i)==' ') str.append("  ");
        }
        //p指向原字符串的末尾
        int p =len-1;
        // q指向现字符串的末尾
        int q=str.length()-1;
        
        while (p>=0) {
            char ch = str.charAt(p--);
            if(ch==' '){
                str.setCharAt(q--, '0');
                str.setCharAt(q--, '2');
                str.setCharAt(q--, '%');
            }else{
                str.setCharAt(q--, ch);
            }
        }

        return str.toString();

    }

    public static void main(String[] args) {
        StringBuffer s = new StringBuffer("tian qi bu cuo ha");
        String s1 = new String();
        ReplaceSpaces solution= new ReplaceSpaces();
        s1 = solution.replaceSpace(s);
        ReplaceSpaces solution2 = new ReplaceSpaces();
        String s2= new String(solution2.replaceSpace1(s));
        System.out.println("这是指针法的结果:"+s2);
        System.out.println("这是遍历法的结果:"+s1);

        
    }
        // replaceSpace.
}
