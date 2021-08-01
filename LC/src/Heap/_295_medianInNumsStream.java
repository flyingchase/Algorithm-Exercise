package Heap;

import java.util.Comparator;
import java.util.PriorityQueue;
import java.util.Queue;

public class _295_medianInNumsStream {
    Queue<Integer> minHeap, maxHeap;

    public void addNum(int num) {
        minHeap = new PriorityQueue<>();
        maxHeap = new PriorityQueue<Integer>(Comparator.reverseOrder());
        if (minHeap.size() == maxHeap.size()) {
            maxHeap.add(num);
            minHeap.add(maxHeap.poll());
        } else {
            minHeap.add(num);
            maxHeap.add(minHeap.poll());
        }
    }

    public double findMedian() {
        return maxHeap.size() == minHeap.size() ? ((minHeap.peek() + maxHeap.peek()) / 2.0) : minHeap.peek();

    }
}
