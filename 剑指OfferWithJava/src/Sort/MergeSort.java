package Sort;

import java.util.Arrays;

@SuppressWarnings({"all"})
public class MergeSort {
    public static void mergeSort(int[] nums) {
        if (nums == null || nums.length <2) return;
        mergeSort(nums, 0, nums.length-1);
    }
    public  static void mergeSort(int[] nums, int l, int r) {
        if(l>=r) return;
//        while (l<r) {

        int mid = l + ((r - l) >> 1);
        mergeSort(nums, l, mid);
        mergeSort(nums, mid + 1, r);
        merge(nums,l,mid,r);
//        }
    }
    public static void merge(int[] nums, int l, int mid,int r) {
        int p1=l,p2=mid+1,i=0;
        int[] helper = new int[r - l + 1];
        while (p1<=mid && p2 <= r ) {
            helper[i++] = nums[p1]<nums[p2] ?nums[p1++] : nums[p2++];
        }
        while (p1 <=mid) {
            helper[i++] = nums[p1++];
        }
        while (p2<=r) {
            helper[i++] = nums[p2++];
        }
        for ( i = 0; i < helper.length; i++) {
            nums[i+l] = helper[i];
        }
    }

    public static void main(String[] args) {
        int[] nums = {1, 0, 9, 8, 7, 2, 3, 6, 5, 4, 13, 11, 12};
        mergeSort(nums);
        System.out.println("Arrays.toString(nums) = " + Arrays.toString(nums));
    }
}



