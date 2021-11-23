package Arrays;

@SuppressWarnings({"all"})
public class replaceBlank05 {
    public static void main(String[] args) {
        StringBuffer str = new StringBuffer("average above and focus it\n");
        System.out.println(replaceBlank(str));
    }

    public static String replaceBlank(StringBuffer str) {
        if(str==null) return null;
        if (str.length()==0||str.indexOf(" ")==-1) {
            return str.toString();
        }

        int blankCount=0;
        for (int i = 0; i < str.length();i++) {
            if(str.charAt(i) ==' ') {
                blankCount++;
            }
        }

        int oringLen= str.length()-1,newLen= str.length()-1+blankCount*2;

        str.setLength(newLen+1);
        while (oringLen>=0&&newLen>=0) {
            if(str.charAt(oringLen)==' ') {
                str.setCharAt(newLen--,'0');
                str.setCharAt(newLen--, '2');
                str.setCharAt(newLen, '%');
            }else {
                str.setCharAt(newLen,str.charAt(oringLen));
            }
            oringLen--;
            newLen--;
        }
        return str.toString();
    }
}
