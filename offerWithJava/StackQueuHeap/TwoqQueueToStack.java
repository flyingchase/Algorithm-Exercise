package StackQueuHeap;

import java.util.LinkedList;
import java.util.Queue;

public class TwoqQueueToStack {
    private Queue<Integer> queue1=new LinkedList<>();
    private Queue<Integer> queue2 = new LinkedList<>();

    public void push(int node) {
        queue1.offer(node);
    }
    public int pop(){
        if(queue1.isEmpty()) {
            throw new RuntimeException("Empty stack!\n");
        }

        while (queue1.size()>1) {
            queue2.offer(queue1.poll());
        }
        int val = queue1.poll();

        Queue<Integer> t =queue1;
        queue1=queue2;
        queue2=t;
        return val;
    }
}
