package Review;

import DataStructure.ListNode;

import java.util.Arrays;

@SuppressWarnings({"all"})
public class sort {

    static int cnt = 0;

    private static void mergesort(int[] arr) {
        if (arr == null) {
            return;
        }
        int[] nums = Arrays.copyOf(arr, arr.length);
        mergesort(nums, 0, nums.length - 1);


    }

    private static void mergesort(int[] nums, int l, int r) {
        if (l == r) {
            return;
        }
        // int cnt = 0;
        cnt++;
        int mid = l + ((r - l) >> 1);
        mergesort(nums, l, mid);
        mergesort(nums, mid + 1, r);
        merge(nums, l, mid, r);

        System.out.println("cnt: " + cnt + Arrays.toString(nums));
    }

    private static void merge(int[] nums, int l, int mid, int r) {
        int p1 = l, p2 = mid + 1, i = 0;
        int[] helper = new int[r - l + 1];

        while (p1 <= mid && p2 <= r) {
            helper[i++] = nums[p1] < nums[p2] ? nums[p2++] : nums[p1++];
        }
        while (p1 <= mid) {
            helper[i++] = nums[p1++];
        }
        while (p2 <= r) {
            helper[i++] = nums[p2++];
        }

        for (int j = 0; j < helper.length; j++) {
            nums[j + l] = helper[j];
        }
    }

    private static void heapsort(int[] nums) {
        if (nums == null) {
            return;
        }
        int[] arr = Arrays.copyOf(nums, nums.length);
        int size = arr.length;
        for (int i = 0; i < size; i++) {
            heapinsert(nums, i);
        }
        while (size-- > 0) {
            swap(nums, 0, size);
            heapify(nums, 0, size);
        }
    }

    private static void heapify(int[] nums, int index, int size) {
        int left = index * 2 + 1;
        while (left < size) {
            int smallest = (left + 1) < size && nums[left] > nums[left + 1] ? left + 1 : left;

            smallest = nums[smallest] < nums[index] ? smallest : index;
            if (smallest == index) {
                break;
            }
            swap(nums, index, smallest);
            index = smallest;
            left = (index * 2) + 1;
        }
    }

    private static void heapinsert(int[] nums, int index) {
        while (nums[index] < nums[(index - 1) / 2]) {

            swap(nums, index, (index - 1) / 2);
            index = (index - 1) / 2;

        }
    }

    private static void quicksort(int[] nums) {
        if (nums == null || nums.length < 2) {
            return;
        }
        int[] arr = Arrays.copyOf(nums, nums.length);
        quicksort(arr, 0, arr.length - 1);
        System.out.println(Arrays.toString(arr));

    }

    private static void quicksort(int[] nums, int l, int r) {
        // if (l == r) return;
        if (l < r) {
            swap(nums, l, r);
            int[] p = paratition(nums, l, r);
            quicksort(nums, l, p[0] - 1);
            quicksort(nums, p[1] + 1, r);
        }
    }

    private static int[] paratition(int[] nums, int l, int r) {
        int less = l - 1, more = r;
        while (l < more) {
            if (nums[l] < nums[r]) {
                swap(nums, ++less, l++);
            } else if (nums[l] > nums[r]) {
                swap(nums, l, --more);
            } else {
                l++;
            }
        }
        swap(nums, more, r);
        return new int[]{less + 1, more};
    }

    private static void swap(int[] nums, int l, int r) {
        int temp = nums[l];
        nums[l] = nums[r];
        nums[r] = temp;

    }

    public static void main(String[] args) {
        int[] nums = {11, 0, 9, 2, 6, 3, 8, 4, 7, 5, 16, 1, 14, 12, 13, 15};

        quicksort(nums);

        // heapsort(nums);
        // System.out.println(Arrays.toString(nums));

        mergesort(nums);
    }

    public ListNode reverse(ListNode root) {
        ListNode prev = null, cur = root;
        while (cur != null) {
            ListNode next = cur.next;
            cur.next = prev;
            prev = cur;
            cur = next;
        }
        return prev;
    }

    

}
