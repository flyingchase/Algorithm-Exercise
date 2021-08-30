package Heap;

import java.util.Arrays;
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

    // use buckrt sort to count the frequency of char
    public String frequencySortWithBuckrtSort (String s) {
        int[][] cnts = new int[128][2];
        char[] ch = s.toCharArray();
        for (int i = 0; i < 128; i++) {
            cnts[i][0]=i;
        }
        for (char c : ch) {
            cnts[c][1]++;
        }
        Arrays.sort(cnts,(a,b)->{
            if (a[1]!=b[1]) {
                return b[1] - a[1];
            }
            return a[0]-b[0];
        });

        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < 128; i++) {
            char c = (char) cnts[i][0];
            int k = cnts[i][1];
            while (k-- > 0) {
                sb.append(c);
            }
        }
        return sb.toString();
    }


}
