package Sort;

import Sort.*;

/* 

*/

public class MergeSort extends AbstractSort {
    private Object[] temp;

    public void sort(Object[] nums) {
        if (nums == null || nums.length <= 1)
            return;

        int n = nums.length;
        temp = new Object[n];
        sortHelper(nums, 0, n - 1);
    }

    private void sortHelper(Object[] data, int p, int r) {
        if (p >= r)
            return;

        // int q = p + (r - p) / 2;
        int q = p +(r-p)/2;
        
        sortHelper(data, p, q);
        sortHelper(data, q + 1, r);
        merge(data, p, q, r);
    }

    private void merge(Object[] data, int p, int q, int r) {
        int i = p, j = q + 1;
        System.arraycopy(data, p, temp, p, r - p + 1);

        for (int k = p; k <= r; k++) {
            if (i > q)
                data[k] = temp[j++];
            else if (j > r)
                data[k] = temp[i++];
            else if (compare(temp[j], temp[i]))
                data[k] = data[j++];
            else
                data[k] = temp[i++];
        }
    }

    public static void main(String[] args) {
        Integer[] nums = new Integer[] { 1, 0, 9, 2, 8, 3, 7, 4, 6, 5, 1, 11, 19, 17, 14 };
        MergeSort mSort = new MergeSort();
        mSort.sort(nums);
        mSort.showArray(nums);

    }
    public static void heapsort(int[] nums) {
        if (nums.length<2||nums==null) {
            return;

        }
        for (int i = 0; i < nums.length; i++) {
            heapinsert(nums,i);
        }
        int size = nums.length;

    }

    public static void heapinsert(int[] nums, int index) {

    }
    public static void heapify(int[] nums, int idnex, int size) {

    }
}
