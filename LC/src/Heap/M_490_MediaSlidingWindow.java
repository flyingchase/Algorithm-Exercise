package Heap;

import java.util.Arrays;
import java.util.Comparator;
import java.util.PriorityQueue;
import java.util.Queue;

/*给你一个数组 nums，有一个长度为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。
你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。*/
public class M_490_MediaSlidingWindow {
    Queue<Integer> maxHeap, minHeap;

    public static void main(String[] args) {
        int[] nums = {11, 0, 9, 2, 6, 3, 8, 4, 7, 5, 16, 1, 14, 12, 13, 15};
        System.out.println(Arrays.toString(new M_490_MediaSlidingWindow().medianSlidingWindow(nums, 4)));
    }
    public double[] medianSlidingWindow(int[] nums, int k) {
        if (nums == null || k < 0 || k > nums.length) {
            return null;
        }
        maxHeap = new PriorityQueue<Integer>(Comparator.reverseOrder());
        minHeap = new PriorityQueue<>();
        double[] res = new double[nums.length - k + 1];
        for (int i = 0; i < k; i++) {
            addNum(nums[i]);
        }
        for (int i = 0, j = i + k; j < nums.length; i++, j++) {
            res[i] = findMedian();
            minHeap.remove(nums[i]);
            maxHeap.remove(nums[j]);
            addNum(nums[j]);

        }

        return res;

    }

    public void addNum(int num) {
        if (maxHeap.size() == minHeap.size()) {
            maxHeap.add(num);
            minHeap.add(maxHeap.peek());
        } else {
            minHeap.add(num);
            maxHeap.add(minHeap.peek());
        }
    }

    public double findMedian() {
        return minHeap.size() == maxHeap.size() ? ((minHeap.peek() + maxHeap.peek()) / 2.0) : (double) minHeap.peek();
    }

}
