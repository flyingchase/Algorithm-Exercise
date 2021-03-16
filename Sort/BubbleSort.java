package Sort;

import java.util.Scanner;

import Sort.*;

/*
冒泡排序:
    时间复杂度:
        - 最好情况 O(N)
        - 最坏情况 O(N^2)
        — 平均情况 O(N^2)
    空间复杂度:
        - 原地排序 O(1)
    稳定排序

*/
public class BubbleSort extends AbstractSort {
    public void sort(Object[] nums) {
        if(nums == null || nums.length<=1) return;

        int n=nums.length;
        //swap标识数组已经排序完成 减少后续操作
        boolean swap = true;
        for(int i=0;swap&&i<n;i++){
            swap=false;
            for(int j=0;j<n-i-1;j++){
                if(compare(nums[j+1], nums[j])) {
                    swap(nums,j,j+1);
                    swap=true;
                }
            }
        }
        
    }

    public static void main(String[] args) {
        Integer[] nums = new Integer[]{1,0,9,2,8,3,7,4,6,5,1,11,19,17,14};
        BubbleSort bubbleSort = new BubbleSort();
        bubbleSort.sort(nums);
        bubbleSort.showArray(nums);
    }
}
