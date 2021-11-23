package StackQueuHeap;

import java.util.LinkedList;
import java.util.Queue;

/* 找出字符流中第一个只出现一次的字符 */
public class FirstAppearOnce {
    // 队列存储字符 
    private int[] cnts = new int[128];
    private Queue<Character> queue = new LinkedList<>();

    public void Insert(char ch) {
        cnts[ch]++;
        queue.add(ch);
        // 若队首字符出现超过一次则剔除
        while (!queue.isEmpty() && cnts[queue.peek()] > 1) {
            queue.poll();
        }
    }

    public char FirstAppearOnceWithQueue() {
        return queue.isEmpty() ? '-1' : queue.peek();
    }
}
