// package Review;
//
// import java.io.BufferedInputStream;
// import java.io.BufferedReader;
// import java.io.InputStreamReader;
// import java.math.BigInteger;
// import java.text.DecimalFormat;
// import java.util.ArrayList;
// import java.util.Arrays;
// import java.util.PriorityQueue;
// import java.util.Scanner;
//
// public class JavaCodesSnippetsDemo {
//
//     public static void main(String args[]) {
//         Scanner cin = new Scanner(new BufferedInputStream(System.in));
// // 2. 读入一个整数
//         int n = cin.nextInt();
// // 3. 读入一个字符串
//         String s = cin.next();
// // 4. 读入一个浮点数
//         double d = cin.nextDouble();
// // 5. 读入一行
//         String s1 = cin.nextLine();
//         // 如果把next（）或者nextInt（），nextDouble() 、 nextFloat()用在nextLine的前面时。nextLine会把前者的结束符“换行符”作为字符串读入，进而不需要从键盘输入字符串nextLine已经转向下一条语句执行
//         // 修正方法：在next()或nextInt()方法使用Enter键之后，填充一个无用的nextLine()
// // 6. 输出
//         System.out.println();
//         System.out.printf("%d %10.5f\n", a, b);
// // 7. 浮点数保留几位小数，DecimalFormat类
// //     0 表示如果位数不足则以 0 填充，# 表示只要有可能就把数字拉上这个位置。这里0指一位数字，#指除0以外的数字。
//         DecimalFormat f = new DecimalFormat("#.##%");
//         DecimalFormat f = new DecimalFormat(",###.##");
//         System.out.println(f.format(d));
// // 8. 大数BigInteger BigDecimal
// //     存储任意精度的数，运算速度比较慢
// //     add,substract,multiply,divide,remainder,compareTo()
// //     divideAndRemainder：a[0]=this / val; a[1]=this % val
// //     pow,gcd,abs,negate,signum,mod,shiftLeft(this<<n),shiftRight(this>>n),and,or,xor
//         BigInteger bi = new BigInteger("2");
//         BigInteger bi = new BigInteger("-101", 2);
//         System.out.println(bi.toString(10));
//         //会使用科学计数法，toPlainString()直接显示
//         // toString()
//         //返回数值上等于此小数，移除末尾的0的BigDecimal
//         // stripTrailingZeros()
// // 9. 进制转换
//         String st = Integer.toString(num, base);
//         int num = Integer.parseInt(st, base);
//
//
//         BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
//         //scanf不能过的一个例子
// //        Scanner cin = new Scanner(new BufferedInputStream(System.in));
//         BufferedReader cin02 = new BufferedReader(new InputStreamReader(System.in));
//         String line;
//         while ((line = cin02.readLine()) != null) {
//             int n = Integer.parseInt(line);
//
//             PriorityQueue<Integer> priorityQueue = new PriorityQueue<>();
//             for (int i = 0; i < n; i++) {
//                 line = cin02.readLine();
//                 String[] items = line.split(" ");
//                 if (Integer.parseInt(items[0]) == 1) {
//                     priorityQueue.offer(Integer.parseInt(items[1]));
//                 } else {
//                     System.out.println(priorityQueue.poll());
//                 }
//             }
//         }
//     }
//
//
//
//
// }
