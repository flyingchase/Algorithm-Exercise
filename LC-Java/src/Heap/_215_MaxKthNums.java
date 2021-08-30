package Heap;

import java.util.Collections;
import java.util.PriorityQueue;

public class _215_MaxKthNums {
    public int MaxKthNums (int[] nums, int k) {
        if (nums==null) {
            return Integer.parseInt(null);
        }

        PriorityQueue<Integer> maxHeap = new PriorityQueue<>(Collections.reverseOrder());
        for (int i = 0; i < nums.length; i++) {
            maxHeap.add(nums[i]);
        }
        while (!maxHeap.isEmpty() && --k > 0) {
            maxHeap.poll();
        }
        return maxHeap.peek();

    }

    public static void main(String[] args) {
        int[] nums = {1,0,9,2,6,3,8,4,7,5,16,1,14,12,13,15};
        System.out.println(new _215_MaxKthNums().MaxKthNums(nums, 3));
    }
}
