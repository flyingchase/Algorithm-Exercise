import java.util.Arrays;

/**
 * SortDailyExercise
 */
public class SortDailyExercise {

    public static void swap(int[] nums, int i ,int j) {
        nums[i] = nums[i] ^ nums[j];
        nums[j] = nums[i] ^ nums[j]; 
        nums[i] = nums[i] ^ nums[j];
    }

    public static void quicksort(int[] nums) {
        if(nums==null||nums.length<2) {
            return;
        }
        quicksort(nums,0,nums.length-1);
    }
    public static void quicksort(int[] nums, int l, int r) {
        if(l>=r) return;
        if(l<r) {
            int[] p =partition(nums, l, r);
            partition(nums, p[1]+1, r);
            partition(nums, l, p[0]-1);
        }
    }

    public static int[]  partition(int[] nums,int l, int r) {
        int less=l-1;
        int more=r;
        while (l<more) {
            if (nums[l]<nums[r]) {
                swap(nums, ++less, l++);
            }else if (nums[l]>nums[r]) {
                swap(nums, --more, l);
            }else{
                l++;
            }
        }
        swap(nums, more, r);
        return new int[] {less+1,more};
    }
    public static void main(String[] args) {
        int[] nums={1,0,9,2,3,6,8,7,5,4};
        quicksort(nums);
        System.out.println("快排: "+Arrays.toString(nums));
    }
}