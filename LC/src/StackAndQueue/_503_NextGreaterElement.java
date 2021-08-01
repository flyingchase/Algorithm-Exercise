package StackAndQueue;

import java.util.Arrays;
import java.util.Stack;

public class _503_NextGreaterElement {
    public int[] nextGreatetElements (int[] nums) {
        if (nums==null) {
            return null;
        }
        int len = nums.length;
        Stack<Integer> stack = new Stack<>();
        int[] res = new int[len];
        Arrays.fill(res, -1);
        for (int i = 0; i < 2 * len; i++) {
            int num = nums[i % len];
            while (!stack.isEmpty() && nums[stack.peek()] < num) {
                res[stack.pop()]=num;
            }
            if (i<len) {
                stack.push(i);
            }
        }
        return res;
    }
}
