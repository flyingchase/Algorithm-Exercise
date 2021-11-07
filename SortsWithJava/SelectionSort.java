package Sort;
import Sort.*;

/* 
选择排序:
    对于一个要排序的数组，我们每次都从数组中寻找最小值，
    并把它和第一个元素交换，然后在剩下的数据中继续寻找最小值，
    然后将其与数组第二个元素交换，如此循环往复，直到整个数组有序

    时间复杂度:
        - 最好最坏平均: O(N^2)
    空间复杂度: O(1)
    不稳定排序 
*/
public class SelectionSort  extends AbstractSort{
    public void sort(Object[] nums) {
        if(nums == null || nums.length<=1) return;

        int n = nums.length;
        for(int i=0;i<n-1;i++) {
            int min=i+1;
            for(int j=i;j<n;j++) {
                if(compare(nums[j],nums[min])) {
                    min=j;
                }
            }
            swap(nums, i, min);
        }
    }

    public static void main(String[] args) {
        Integer[] nums = new Integer[]{1,0,9,2,8,3,7,4,6,5,1,11,19,17,14};
        SelectionSort sSort = new SelectionSort();
        sSort.sort(nums);
        sSort.showArray(nums);

    }
}
