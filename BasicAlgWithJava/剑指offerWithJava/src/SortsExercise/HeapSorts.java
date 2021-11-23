package SortsExercise;

import java.util.Arrays;

import static DataStructure.swap.swap;

@SuppressWarnings({"all"})
public class HeapSorts {
    public static void main(String[] args) {

        int[] nums= {1,0,9,7,2,3,5,6,4,13,14,12,11};

        heapsort(nums);
        System.out.println(Arrays.toString(nums));
    }
    public static void heapsort(int[] nums) {

        if (nums.length<2||nums==null) {
            return;
        }
        int size = nums.length;
        for (int i = 0; i < nums.length; i++) {
            heapinsert(nums,i);
        }

        while (size > 0) {
            swap(nums,0,--size);
            heapify(nums,0,size);
        }
    }
    public static void heapinsert(int[] nums,int index) {
        int parent = (index - 1) / 2;
        while (nums[index] > nums[parent]) {
            swap(nums, index, parent);
            index=parent;
        }
    }

    public static void heapify(int[] nums, int index, int size) {
        int left = 2 * index + 1;
        int right = 2 * index + 2;
        while (left<size) {
            int largest = (right<size&&nums[left]<nums[right])?right:left;
            largest = nums[index] > nums[largest] ? index : largest;
            if (largest==index) {
                break;
            }
            swap(nums,index,largest);
            index=largest;
            left=2*index+1;
            right=2*index + 2;
        }
    }



}
