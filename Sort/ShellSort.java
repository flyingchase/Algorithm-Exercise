package Sort;
import Sort.*;

/* 
希尔排序: (插入排序的优化版本)
    数据分成了 n 个组，对每个组内的数据分别进行插入排序，
    这样就能够将较大的数据一次性移动很多位，为后续的比较交换提供了便利。

    

*/

public class ShellSort extends AbstractSort {
    public void sort(Object[] nums) {
        if(nums == null || nums.length<=1) return;

        int n= nums.length;
        int k = n/2;
        while (k>=1) {
            for(int i=0;i<n-k;i++) {
                int j = i+k;
                Object v = nums[j];
                while (j>k-1 && compare(v, nums[j-k])) {
                    nums[j]=nums[j-k];
                    j-=k;
                }
                nums[j]=v;
            }
            k/=2;
        }
    }

    public static void main(String[] args) {
        Integer[] nums = new Integer[]{1,0,2,9,3,8,4,7,5,6};
        ShellSort ssSort = new ShellSort();
        ssSort.sort(nums);
        ssSort.showArray(nums);
    }
}
