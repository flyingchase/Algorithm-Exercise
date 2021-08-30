package String;

/*给定一个字符串 s 和一个整数 k，从字符串开头算起，每 2k 个字符反转前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。*/
public class E_541_reverseStr {
    public static void main(String[] args) {
        String s = "abcd";
        System.out.println(new E_541_reverseStr().reverseStr(s, 2));
    }

    // public String reverseStr(String s, int k) {
    //     if (s == null) {
    //         return s;
    //     }
    //     char[] ch = s.toCharArray();
    //     int len = ch.length, count = 0;
    //     boolean Kflag = false;
    //     for (int i = 0; i < len; i++) {
    //         int left = len - i;
    //         if (left < k) {
    //             reverse(ch, i, len-1);
    //             break;
    //         } else if (left < 2 * k) {
    //             reverse(ch, i, i + k - 1);
    //             break;
    //         } else {
    //             if (i % k == 0 && i != 0) {
    //                 count++;
    //                 Kflag = !Kflag;
    //             }
    //             if (Kflag) {
    //                 reverse(ch, (count - 1) * k, i);
    //             }
    //         }
    //     }
    //     return String.valueOf(ch);
    // }

    public String reverseStr(String s, int k) {
        char[] arr = s.toCharArray();
        int n = arr.length;
        int i = 0;
        while(i < n) {
            int j = Math.min(i + k - 1, n - 1);
            swap(arr, i, j);
            i += 2 * k;
        }
        return String.valueOf(arr);
    }
    private void swap(char[] arr, int l, int r) {
        while (l < r) {
            char temp = arr[l];
            arr[l++] = arr[r];
            arr[r--] = temp;
        }
    }

    private void reverse(char[] ch, int i, int j) {
        for (; i < j; i++, j--) {
            char temp = ch[i];
            ch[i] = ch[j];
            ch[j] = temp;
        }
    }
}
