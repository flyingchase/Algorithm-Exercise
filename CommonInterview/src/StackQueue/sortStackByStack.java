package StackQueue;

import java.util.Stack;
@SuppressWarnings({"all"})
public class sortStackByStack {

    public static void sortStackByStack(Stack<Integer> stack) {
        Stack<Integer> help = new Stack<Integer>();
        if (!stack.isEmpty()) {
            int cur = stack.pop();
            while (!help.isEmpty() && help.peek() < cur) {
                stack.push(help.pop());
            }
            help.push(cur);
        }
        if (!help.isEmpty()) {
            stack.push(help.pop());
        }

    }
}
