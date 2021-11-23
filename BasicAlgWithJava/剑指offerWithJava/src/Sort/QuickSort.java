package Sort;

import java.util.Arrays;

import static DataStructure.swap.swap;

@SuppressWarnings({"all"})
public class QuickSort {
    public static void quickSort(int[] nums) {
        if (nums == null || nums.length < 2) {
            return;
        }
        quickSort(nums, 0, nums.length - 1);
    }
    public static void quickSort(int[] nums, int l, int r) {
        if(l < r) {
            swap(nums, l, r);
            int[] partitions = partitions(nums, l, r);

            partitions(nums, l, partitions[0]-1);
            partitions(nums,partitions[1]+1, r);
        }
    }
    public static int[] partitions(int[] nums, int l, int r) {
        int less=l,more = r;
        while (l<more) {
            if (nums[l]<nums[r]) {
                swap(nums, less++,l++);
            }else if (nums[l]> nums[r]) {
                swap(nums, l, --more);
            }else {
                l++;
            }

        }
        swap(nums, more, r);
        return new int[]{less , more};

    }

    public static void main(String[] args) {
        int[] nums= {1,0,9,8,7,2,3,6,5,4,13,11,12};
        quickSort(nums);
        System.out.println("Arrays.toString(nums) = " + Arrays.toString(nums));
    }
}
