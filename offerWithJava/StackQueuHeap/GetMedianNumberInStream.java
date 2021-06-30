package StackQueuHeap;

import java.util.Comparator;
import java.util.PriorityQueue;

/* 数据流中的中位数 */
/* 利用大根堆存储较小的一半元素,小根堆存储较大的一半元素并维持两堆size差值在1以内 */
/* 使得中位数在大根堆和小根堆堆顶产生 */
public class GetMedianNumberInStream {
    // 优先队列默认为小根堆 故使用Comparator.reverseOrder()
    private PriorityQueue<Integer> maxHeap = new PriorityQueue<>(Comparator.reverseOrder());
    private PriorityQueue<Integer> minHeap = new PriorityQueue<>();

    public void Insert(Integer num) {
        // 大根堆为空或者读入数字小于大根堆堆顶 入大根堆利用特点使得堆顶为最大值
        if (maxHeap.isEmpty() || num < maxHeap.peek()) {
            maxHeap.offer(num);
            // 大根堆size大于小根堆则将大根堆堆顶元素入小根堆 使得小根堆内元素为较大部分
            if (maxHeap.size() - minHeap.size() > 1) {
                minHeap.offer(maxHeap.poll());
            }
        } else {
            minHeap.offer(num);
            if (minHeap.size() - maxHeap.size() > 1) {
                maxHeap.offer(maxHeap.poll());
            }
        }

    }

    public Double GetMediaNubmer() {
        if (maxHeap.size() > minHeap.size()) {
            return (double) maxHeap.peek();
        } else {
            return (double) (maxHeap.size() < minHeap.size() ? minHeap.peek() : ((maxHeap.peek() + minHeap.peek()) / 2.0));
        }
    }
}


// lambda表达式建大根堆
class GetMedianNumberInStream02 {
    private PriorityQueue<Integer> maxHeap = new PriorityQueue<>((o1, o2) -> o2 - o1);
    private PriorityQueue<Integer> minHeap = new PriorityQueue<>();
    // 数据流读入的元素个数
    private N=0;

    public void Insert02(Integer val) {
        if (N % 2 == 0) {
            /* N为偶数插入minHeap */
            /* 保证minHeap内元素为较大部分 故先进入大根堆取出堆顶元素再插入小根堆 */
            maxHeap.add(val);
            minHeap.add(maxHeap.poll());
        } else {
            /* N为奇数插入maxHeap */
            /* 保证maxHeap为较小部分 先入小根堆再取出小根堆堆顶元素插入大根堆 */
            minHeap.add(val);
            maxHeap.add(minHeap.poll());
        }
        N++;
    }

    public Double getMedia() {
        if (N % 2 == 0) {
            return ((minHeap.peek() + maxHeap.peek()) / 2.0)
        } else {
            return (double) (minHeap.peek());
        }

    }

}
