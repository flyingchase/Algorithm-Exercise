package StackQueuHeap;

import java.util.LinkedList;
import java.util.Queue;

public class TwoqQueueToStack {
    private Queue<Integer> queue1=new LinkedList<>();
    private Queue<Integer> queue2 = new LinkedList<>();

    public void push(int node) {
        queue1.offer(node);
    }
}
