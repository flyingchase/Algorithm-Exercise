package Heap;

import java.lang.invoke.VarHandle;
import java.util.Arrays;
import java.util.PriorityQueue;

public class _480_medianSlidingWindow {
    public static void main(String[] args) {
        int[] nums = {1, 3, -1, -3, 5, 3, 6, 7};
        int[] ints = {2147483647, 2147483647};
        System.out.println(new _480_medianSlidingWindow().findMedianInNums(ints));
        System.out.println(Arrays.toString(new _480_medianSlidingWindow().medianSlidingWindow(nums, 3)));
    }
    public double[] medianSlidingWindow(int[] nums, int k) {
        if (nums == null || k < 0 || k > nums.length) {
            return null;
        }
        double[] res = new double[nums.length - k + 1];
        int[] windows = Arrays.copyOfRange(nums,0,k);
        res[0]=(findMedianInNums(windows));
        int j=1;

        for (int i = k; i < nums.length; i++) {
            for (int n = 0; n < windows.length-1; n++) {
                windows[n]=windows[n+1];
            }
            windows[k-1]=nums[i];
            res[j++]=findMedianInNums(windows);
        }
        return res;
    }


    private double findMedianInNums (int[] nums) {
        PriorityQueue<Integer> minHeap = new PriorityQueue<>();
        PriorityQueue<Integer> maxHeap = new PriorityQueue<>((a, b) -> (Integer.compare(a,b)));

        for (int num : nums) {
            if (maxHeap.size()==minHeap.size()) {
                maxHeap.offer(num);
                minHeap.offer(maxHeap.poll());
            }else {
                minHeap.offer(num);
                maxHeap.offer(minHeap.poll());
            }
        }

        if (minHeap.size()==maxHeap.size()) {
            return ((double)minHeap.poll()+maxHeap.poll())/2.0;
        }else {
            return minHeap.poll()*1.0;
        }
    }

}
