package Array;

import java.util.PriorityQueue;

/*给定一个整数数组 arr，如果它是有效的山脉数组就返回 true，否则返回 false。*/
public class E_941_validMountainArray {

    public static void main(String[] args) {
        int[] arr = new int[]{0, 3, 2, 1};
    }

    public boolean validMountainArray(int[] arr) {
        boolean flag = false;
        PriorityQueue<Integer> maxHeap = new PriorityQueue<>();
        for (int i : arr) {
            maxHeap.add(i);
        }
        Integer maxInteger = maxHeap.poll();
        int max = maxInteger.intValue();
        int index = Integer.MIN_VALUE;
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] == max) {
                index = i;
                break;
            }
        }



        int l = 0, r = arr.length - 1, count = 0;
        while (l <= r && l < arr.length - 1 && r > 0) {
            if (arr[l] < arr[l + 1] && arr[r] < arr[r - 1]) {
                count++;
            }
        }
        if (count == (arr.length / 2)) {
            return true;
        }
        return flag;

    }


}
