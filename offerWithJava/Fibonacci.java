package 剑指offer_Java;

import java.util.Scanner;

public class Fibonacci {
    public static void main(String[] args) {
        Fibonacci f = new Fibonacci();
        Scanner in = new Scanner(System.in);
        int i = in.nextInt();
        System.out.println(f.Fibonacci1(i));
    }

    public int Fibonacci1(int n) {
        if (n < 2) return n;
        int[] res = new int[n + 1];
        res[0] = 0;
        res[1] = 1;
        for (int i = 2; i <= n; ++i) {
            res[i] = res[i - 1] + res[i - 2];
        }
        return res[n];
    }
}
