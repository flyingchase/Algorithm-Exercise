package 剑指offer_Java;

import java.util.Scanner;

import 剑指offer_Java.*;

/*
输入数字 n，按顺序打印出从 1 最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数即 999。


*/
public class PrintToMaxN {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        PrintToMaxN maxN = new PrintToMaxN();
        maxN.printToMaxN1(n);
    }

    // 数字转化成数组
    private void PrintNumber(char[] chars) {
        int index = 0;
        int n = chars.length;
        for (char ch : chars) {
            if (ch != '0') break;
            ++index;
        }
        StringBuilder sb = new StringBuilder();
        for (int i = index; i < n; ++i) {
            sb.append(chars[i]);
        }
        System.out.println(sb.toString());
    }

    // 检查是否溢出
    private boolean increment(char[] chars) {
        boolean flag = false;
        int n = chars.length;
        int carry = 1;
        for (int i = n - 1; i >= 0; --i) {
            int num = chars[i] - '0' + carry;
            if (num > 9) {
                if (i == 0) {
                    flag = true;
                    break;
                }
                chars[i] = '0';
            } else {
                ++chars[i];
                break;
            }
        }
        return flag;
    }

    // 输出打印
    public void printToMaxN1(int n) {
        if (n < 1)
            return;
        char[] chars = new char[n];
        for (int i = 0; i < n; ++i) {
            chars[i] = '0';
        }
        while (!increment(chars)) {
            PrintNumber(chars);

        }
    }

}
