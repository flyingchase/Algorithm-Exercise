package Heap;

import java.util.*;
import java.util.stream.Collectors;

public class Lint_604_WinSum {
    public static void main(String[] args) {
        int[] nums = {11, 0, 9, 2, 6, 3, 8, 4, 7, 5, 16, 1, 14, 12, 13, 15};
        System.out.println(Arrays.toString(new Lint_604_WinSum().winSum(nums,3)));
    }

    public int[] winSum(int[] nums, int k) {
        if (nums == null || k > nums.length) {
            return null;
        }
        // int[] res = new int[nums.length-k+1];
        ArrayList<Integer> integers = new ArrayList<>();
        PriorityQueue<Integer> windows = new PriorityQueue<>((o1, o2) -> (o2 - o1));
        for (int i = 0; i < k; i++) {
            windows.add(nums[i]);
        }
        // res[0]= windows.peek();

        integers.add(windows.peek());
        for (int i = 0, j = i + k; j < nums.length; i++, j++) {
            windows.remove(nums[i]);
            windows.add(nums[j]);
            // res[i+1]=windows.poll();
            integers.add(windows.peek());
        }

        int[] res = integers.stream().mapToInt(Integer::intValue).toArray();


        return res;

    }
}
