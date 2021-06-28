package Sort;

import java.util.Arrays;

import static DataStructure.swap.swap;
//import static com.sun.tools.javac.jvm.ByteCodes.swap;

@SuppressWarnings({"all"})
public class HeapSort02 {
    public static void main(String[] args) {
        int[] nums =         {1,0,9,8,7,2,3,6,5,4,13,11,12};
        heapsort(nums);
        System.out.println(Arrays.toString(nums));

    }
    public static void heapsort(int[] nums) {
        if (nums.length <2 || nums==null) {
            return;
        }
        for (int i = 0; i < nums.length; i++) {
            heapinsert(nums,i);
        }
        int size = nums.length;
        while (size > 0) {
            swap(nums, 0, --size);
            heapify(nums, 0,size);
        }
    }

    public static void heapinsert(int[] nums, int index) {
        while (nums[index]>nums[(index-1)/2]) {
            swap(nums, index, (index-1)/2);
            index=(index - 1)/2;
        }
    }

    public static void heapify(int[] nums,int index, int size) {
        int left = 2*index+1;
        while (left < size) {
            int largest = (left + 1<size)&&(nums[left+1]>nums[left])?left+1:left;
            largest=nums[largest]>nums[index]?largest:index;
            if (index==largest) {
                break;
            }
            swap(nums, index, largest);
            index=largest;
            left=2*index + 1;
        }
    }

}
