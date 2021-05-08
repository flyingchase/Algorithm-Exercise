import java.util.Arrays;

/*
 * @Author: qlzhou  
 * @Date:   2021-05-04 15:00:12
 * @Last Modified by:   gunjianpan
 * @Last Modified time: 2021-05-05 16:59:44
 */

public class SortExercise {
    public static void swap(int[] nums, int i, int j) {
        nums[i] = nums[i] ^ nums[j];
        nums[j] = nums[i] ^ nums[j];
        nums[i] = nums[i] ^ nums[j];
    }

    public static void mergeSort(int[] nums) {
        if (nums == null || nums.length < 2) {
            return;
        }
        mergeSort(nums, 0, nums.length - 1);
    }

    public static void mergeSort(int[] nums, int l, int r) {
        if (l <= r)
            return;
        int mid = l + ((r - l) >> 1);
        mergeSort(nums, l, mid);
        mergeSort(nums, mid + 1, l);
        merge(nums, l, mid, r);
    }

    public static void merge(int[] nums, int l, int mid, int r) {
        int p1 = l, p2 = mid + 1, i = 0;
        int[] helper = new int[r - l + 1];
        while (p1 <= mid && p2 <= r) {
            helper[i++] = nums[p1] < nums[p2] ? nums[p1++] : nums[p2++];
        }
        while (p1 <= mid) {
            helper[i++] = nums[p1++];
        }
        while (p2 <= r) {
            helper[i++] = nums[p2++];
        }
        for (int k = 0; k < helper.length; k++) {
            nums[k + l] = helper[k];
        }
    }

    public static void heapsort(int[] nums) {
    
        if (nums==null||nums.length<2){
            return;
        }
        for (int i = 0; i < nums.length; i++) {
            heapinsert(nums, i);
        }
        int size = nums.length;
        
        while (size>0) {
            swap(nums, 0, --size);
            heapify(nums, 0, size);
        }
        
    }
    public static void heapify(int[] nums, int index, int size) {
        int left = (2*index+1);
        while (left<size) {
            int largest = (left+1<size&&nums[left+1]>nums[left])?left+1:left;
            largest = nums[index]>nums[largest]?index:largest;
            if(largest==index) {
                break;
            }
            swap(nums, index, largest);
            index=largest;
            left=2*index+1;
        }

    }
    public static void heapinsert(int[] nums, int index) {
        int parents= (index-1)/2;
        while (nums[index]>nums[parents]) {
            swap(nums, index, parents);
            index=parents;
        }
    }

    public static void quicksort(int[] nums) {
        if(nums==null||nums.length<2) {
            return;
        }

        quicksort(nums,0,nums.length-1);
    }
    public static void quicksort(int[] nums, int l, int r) {
        if(l==r) return;
        if(l<r) {
            int 
        }
    }
    public static int[] partition(int[] nums, int l, int r) {
        
    }
    
    public static void main(String[] args) {
        int[] nums = { 1, 0, 9, 2, 3, 8, 4, 6, 7, 5 };
        int[] arr = Arrays.copyOf(nums, nums.length);
        // mergeSort(nums);
        heapsort(nums);
        System.out.println(Arrays.toString(arr));
        System.out.println(Arrays.toString(nums));
    }

    
}
