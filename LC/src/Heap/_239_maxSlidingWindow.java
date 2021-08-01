package Heap;

import java.util.*;

/*给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

*/
public class _239_maxSlidingWindow {

    int[] nums = {2, 1, 2, 4, 3};

    public static void main(String[] args) {
        int[] nums = {11, 0, 9, 2, 6, 3, 8, 4, 7, 5, 16, 1, 14, 12, 13, 15};
        System.out.println(Arrays.toString(new _239_maxSlidingWindow().maxSlidingWindows(nums, 3)));
    }

    private static int[] nextGreaterElemenet(int[] nums) {
        int[] res = new int[nums.length];
        Stack<Integer> stack = new Stack<>();
        for (int i = nums.length - 1; i >= 0; i--) {
            // 维护单调栈的递增顺序 当要加入的值大于栈顶元素是 抛出栈顶丢弃
            while (!stack.isEmpty() && nums[i] >= stack.peek()) {
                stack.pop();
            }
            // add the result to stack
            res[i] = stack.isEmpty() ? -1 : stack.peek();
            // action to come in stack
            stack.push(nums[i]);
        }
        return res;
    }

    // TreeSet

    public int[] maxSlidingWindows(int[] nums, int k) {
        if (nums == null || k < 0 || k > nums.length) {
            return null;
        }
        // ArrayList<Integer> arr = new ArrayList<>();
        PriorityQueue<Integer> maxHeap = new PriorityQueue<>(Comparator.reverseOrder());
        int[] res02 = new int[nums.length - k + 1];
        for (int i = 0; i < k; i++) {
            maxHeap.add(nums[i]);
        }

        // arr.add(maxHeap.peek());
        res02[0] = maxHeap.peek();
        for (int i = 0, j = i + k; j < nums.length; j++, i++) {
            maxHeap.remove(nums[i]);
            maxHeap.add(nums[j]);
            // arr.add(maxHeap.peek());
            res02[i + 1] = maxHeap.peek();
        }
        // int[] res = arr.stream().mapToInt(Integer::intValue).toArray();
        // System.out.println(Arrays.toString(res02));
        return res02;

    }

    // O(N) 单调栈 方式的时间复杂度为
    public int[] maxSlidingWindowWithDqueue(int[] nums, int k) {
        Deque<Integer> queue = new ArrayDeque<>();
        int index = 0;
        int[] res = new int[nums.length - k + 1];
        for (int i = 0; i < nums.length; i++) {
            while (!queue.isEmpty() && nums[i] > queue.peekLast()) {
                queue.pollLast();
            }
            queue.addLast(nums[i]);
            if (i >= k - 1) {
                res[index++] = queue.peekFirst();
                if (i >= k - 1 && nums[i - k + 1] == queue.peekFirst()) {
                    queue.removeFirst();
                }
            }
        }
        return res;

    }

}


























