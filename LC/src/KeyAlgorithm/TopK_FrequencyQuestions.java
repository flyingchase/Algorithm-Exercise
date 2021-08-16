package KeyAlgorithm;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.PriorityQueue;

/*
 * TopK Frequent elements
 * */
public class TopK_FrequencyQuestions {

    private List<Integer> topKFrequencts(int[] nums, int k) {
        PriorityQueue<Node> maxHeap = new PriorityQueue<>();
        HashMap<Integer, Node> map = new HashMap<>();
        List<Integer> list = new ArrayList<>();
        for (int i = 0; i < nums.length; i++) {
            if (!map.containsKey(nums[i])) {
                map.put(nums[i], new Node(nums[i], 1));
            } else {
                Node node = map.get(nums[i]);
                node.frequency++;
                map.put(nums[i], node);
            }
        }

        for (Integer key : map.keySet()) {
            maxHeap.offer(map.get(key));
        }
        // maxHeap.addAll(map.values());
        for (int i = 0; i < k; i++) {
            list.add(maxHeap.poll().frequency);
        }
        return list;

    }

    private class Node implements Comparable<Node> {
        int key;
        int frequency;

        public Node(int key, int frequency) {
            this.key = key;
            this.frequency = frequency;
        }

        @Override
        public int compareTo(Node o) {
            return o.frequency - this.frequency;
        }
    }
}
