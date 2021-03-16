package Sort;
import Sort.*;
/* 
插入排序:
    依次遍历每一个数字，并将其和前面已排序的数据进行对比，将其插入到合适的位置。
    时间复杂度:
        - O(N^2)
    空间复杂度:
        - O(1)
    稳定排序
*/

public class InsertionSort extends AbstractSort {
    public void sort(Object[] nums) {
        if(nums == null || nums.length<=1) return;

        int n = nums.length;
        for(int i=0;i<n-1;i++) {
            int j=i+1;
            //k暂存待插入位置的数nums[j]
            Object k = nums[j];
            //寻找到合适的插入位置
            while (j>0 && compare(k, nums[j-1])) {
                nums[j] = nums[--j];
            }
            //插入到比之前的nums[j]还小的位置
            nums[j] = k;
        }
    }

    public static void main(String[] args) {
        Integer[] nums = new Integer[] { 1, 0, 9, 2, 8, 3, 7, 4, 6, 5, 1, 11, 19, 17, 14 };
        InsertionSort iSort = new InsertionSort();
        iSort.sort(nums);
        iSort.showArray(nums);
    }
}
