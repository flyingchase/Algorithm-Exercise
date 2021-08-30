package StackAndQueue;

import java.util.HashMap;
import java.util.Map;
import java.util.Stack;

/*
* 给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。

请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
* */
public class _496_NextGreaterElement {

    // monotonic stack 单调栈
    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        Map<Integer, Integer> map = new HashMap<>();
        Stack<Integer> stack = new Stack<>();
        for (int i = nums2.length - 1; i >= 0; i--) {
            while (!stack.isEmpty() && nums2[i] >= stack.peek()) {
                stack.pop();
            }
            map.put(nums2[i], stack.isEmpty() ? -1 : stack.peek());
            stack.push(nums2[i]);
        }
        int[] res = new int[nums1.length];
        for (int i = 0; i < nums1.length; i++) {
            res[i] = map.get(nums1[i]);
        }
        return res;
    }

    public int[] nextGreaterElement02(int[] nums1, int[] nums2) {
        Map<Integer, Integer> map = new HashMap<>();
        int[] res = new int[nums1.length];
        Stack<Integer> integers = new Stack<>();
        for (int num : nums2) {
            while (!integers.isEmpty() && integers.peek() < num) {
                map.put(integers.pop(),num);
            }
            integers.push(num);
        }
        for (int i = 0; i < nums1.length; i++) {
            res[i]=map.getOrDefault(nums1[i],-1);
        }
        return res;

    }

}
