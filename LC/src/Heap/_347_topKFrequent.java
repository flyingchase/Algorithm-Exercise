package Heap;

import java.util.*;

/*
*
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。
 */
public class _347_topKFrequent {

    public static void main(String[] args) {
        int[] nums = {1, 1, 1, 2, 2, 3};
        System.out.println(Arrays.toString(new _347_topKFrequent().topKFrequent(nums, 2)));
    }
    public int[] topKFrequent(int[] nums, int k) {
        int[] result = new int[k];
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int num : nums) {
            map.put(num, map.getOrDefault(num, 0) + 1);
        }

        Set<Map.Entry<Integer, Integer>> entries = map.entrySet();
        // 根据map的value值正序排，相当于一个小顶堆
        PriorityQueue<Map.Entry<Integer, Integer>> queue = new PriorityQueue<>((o1, o2) -> o1.getValue() - o2.getValue());
        for (Map.Entry<Integer, Integer> entry : entries) {
            queue.offer(entry);
            if (queue.size() > k) {
                queue.poll();
            }
        }
        for (int i = k - 1; i >= 0; i--) {
            result[i] = queue.poll().getKey();
        }
        return result;

        // Arrays.stream(objects).map(Integer::intValue).toArray()
    }
}
