package StackQueuHeap;

import java.util.Stack;

/* 
栈的弹出顺序是否符合规定

使用一个栈来模拟压入弹出操作。每次入栈一个元素后，都要判断一下栈顶元素是不是当前出栈序列 popSequence 的第一个元素，如果是的话则执行出栈操作并将 popSequence 往后移一位，继续进行判断。


*/
public class IsPopOrder {
    public static boolean isPopOrder(int[] pushSequence, int[] popSequence) {
        int n = pushSequence.length;
        Stack<Integer> stack = new Stack<>();
        for (int pushIndex = 0, popIndex = 0; pushIndex < n; pushIndex++) {
            stack.push(pushSequence[pushIndex]);
            while (popIndex < n && !stack.isEmpty() && stack.peek() == popSequence[popIndex]) {
                stack.pop();
                popIndex++;
            }
        }
        return stack.isEmpty();
    }

    public static void main(String[] args) {
        int[] popSequence = {1, 2, 3, 4, 5, 6, 7, 8, 9, 0};
        int[] pushSequence = {0, 9, 8, 7, 6, 5, 4, 3, 2, 1};
        System.out.println((isPopOrder(pushSequence, popSequence)) ? "yes" : "no");
    }
}
