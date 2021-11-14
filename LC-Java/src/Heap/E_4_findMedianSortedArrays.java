package Heap;

import java.util.Comparator;
import java.util.PriorityQueue;


public class E_4_findMedianSortedArrays {
    PriorityQueue<Integer> minHeap, maxHeap;

    public static void main(String[] args) {
        int[] nums1 = {1, 2}, nums2 = {3, 4};
        System.out.println(new E_4_findMedianSortedArrays().findMedianSortedArrays(nums1, nums2));

    }

    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        if (nums1 == null && nums2 == null) {
            return Double.parseDouble(null);
        }
        int m = nums1.length;
        int n = nums2.length;
        minHeap = new PriorityQueue<Integer>();
        maxHeap = new PriorityQueue<Integer>(Comparator.reverseOrder());
        for (int i = 0; i < nums1.length; i++) {
            addNum(nums1[i]);
        }
        for (int i = 0; i < nums2.length; i++) {
            addNum(nums2[i]);
        }
        return findMedian();
    }

    private void addNum(int num) {
        if (minHeap.size() == maxHeap.size()) {
            maxHeap.add(num);
            minHeap.add(maxHeap.poll());
        } else {
            minHeap.add(num);
            maxHeap.add(minHeap.poll());
        }
    }

    private double findMedian() {
        double v = minHeap.size() == maxHeap.size() ? (minHeap.peek() + maxHeap.peek()) / 2.0 : (double) (minHeap.peek());
        return v;
    }
}
