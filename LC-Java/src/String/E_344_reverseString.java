package String;

public class E_344_reverseString {
    public static void main(String[] args) {
        char[] s = {};
        new E_344_reverseString().reverseString(s);
        System.out.println(s);
    }
    public void reverseString(char[] s) {
        if (s==null||s.length<2) {
            return;
        }
        for (int i = 0, j = s.length - 1; i < j; i++, j--) {
            swap(s, i, j);
        }
    }

    private void swap(char[] s, int start, int end) {
        char temp = s[start];
        s[start] = s[end];
        s[end] = temp;
    }
}
