package Sort;

import java.util.Arrays;

import static DataStructure.swap.swap;

public class BubbleSort {
    public static  void bubbleSort(int[] nums) {
       if (nums == null || nums.length <= 2) return;
       bubbleSort(nums, 0, nums.length-1);
    }
    public static void bubbleSort(int[] nums, int left, int right) {
        for (int i = 0; i < nums.length-1; i++) {
            for (int j = 0; j < nums.length-1-i; j++) {
                if (nums[j]>nums[j+1]) {
                    swap(nums, j+1, j);
                }
            }
        }
    }
    public static void main(String[] args) {
        int[] nums = {1,0,9,8,7,2,3,6,5,4,13,11,12};
        bubbleSort(nums);
        System.out.println("Arrays.toString(nums) = " + Arrays.toString(nums));
    }
}
