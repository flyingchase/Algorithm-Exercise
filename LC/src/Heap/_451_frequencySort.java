package Heap;

import java.util.HashMap;
import java.util.Map;
import java.util.PriorityQueue;

public class _451_frequencySort {


    public static void main(String[] args) {
        String s = "woaiwenleiwenleiwoshizhuwoshizhunizhidaoma";
        System.out.println(new _451_frequencySort().frequencySort(s));
    }
    public String frequencySort(String s) {
        if (s == null) {
            return null;
        }

        PriorityQueue<Map.Entry<Character, Integer>> pq = new PriorityQueue<>((a, b) -> (b.getValue() - a.getValue()));

        Map<Character, Integer> map = new HashMap<>();
        char[] ch = s.toCharArray();
        for (char c : ch) {
            map.put(c, map.getOrDefault(c, 0) + 1);
        }

        pq.addAll(map.entrySet());
        StringBuilder sb = new StringBuilder();
        while (!pq.isEmpty()) {
            Map.Entry<Character, Integer> poll = pq.poll();
            for (int i = 0; i < poll.getValue().intValue(); i++) {
                sb.append(poll.getKey());
            }
        }

        return sb.toString();
    }


}
