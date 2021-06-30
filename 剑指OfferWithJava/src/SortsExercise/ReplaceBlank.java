package SortsExercise;
@SuppressWarnings({"all"})
public class HeapSort02 {
    public static void main(String[] args) {
        int[] nums = new int[] {
            1,2,0,9,3,4,8,7,6,11,5
        };

        StringBuffer str = new StringBuffer("average above and focus it\n");
        System.out.println(replaceBlank(str));
    }

    public static void heapsort(int[] nums) {
        if (nums==null||nums.length <2) {

            return;
        }
    }

    public static void heapsort(int[] nums, int l, int r) {

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
