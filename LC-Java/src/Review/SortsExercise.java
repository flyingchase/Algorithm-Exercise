package Review;

import DataStructure.ListNode;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Set;

public class SortsExercise {
    public static void main(String[] args) {
        int[] nums = {11, 0, 9, 2, 6, 3, 8, 4, 7, 5, 16, 1, 14, 12, 13, 15};
        int[] quicksort = new SortsExercise().quicksort(nums);
        System.out.println("quicksort  result = " + Arrays.toString(quicksort));
        // System.out.printf("quicksort result: ", Arrays.toString(quicksort));
        int[] heapsort = new SortsExercise().heapsort(nums);
        System.out.println("heapsort result = " + Arrays.toString(heapsort));
    }

    public void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

    public int[] quicksort(int[] arr) {
        if (arr == null || arr.length < 2) {
            return arr;
        }
        int[] nums = Arrays.copyOf(arr, arr.length);
        quicksort(nums, 0, nums.length - 1);


        return nums;
    }

    private void quicksort(int[] nums, int l, int r) {
        if (l < r) {
            swap(nums, l, r);
            int[] p = paratition(nums, l, r);
            quicksort(nums, p[1] + 1, r);
            quicksort(nums, l, p[0] - 1);
        }
    }

    private int[] paratition(int[] nums, int l, int r) {
        int less = l - 1, more = r;
        while (l < more) {
            if (nums[l] > nums[r]) {
                swap(nums, ++less, l++);
            } else if (nums[l] < nums[r]) {
                swap(nums, --more, l);
            } else {
                l++;
            }
        }
        swap(nums, more, r);
        return new int[]{less + 1, more};
    }

    public int[] heapsort(int[] arr) {
        if (arr == null || arr.length < 2) {
            return arr;
        }
        int[] nums = Arrays.copyOf(arr, arr.length);
        heapsort(nums, 0, nums.length - 1);
        return nums;
    }

    private void heapsort(int[] nums, int l, int r) {
        if (l == r) {
            return;
        }
        for (int i = 0; i < nums.length; i++) {
            heapInsert(nums, i);
        }
        int size = nums.length;
        while (size-- > 0) {
            swap(nums, 0, size);
            heapIfy(nums, 0, size);
        }
    }

    private void heapIfy(int[] nums, int index, int size) {
        int left = index * 2 + 1;
        while (left < size) {
            int smallest = left + 1 < size && nums[left + 1] < nums[left] ? left + 1 : left;
            smallest = nums[index] < nums[smallest] ? index : smallest;
            if (smallest == index) {
                break;
            }
            swap(nums, index, smallest);
            index = smallest;
            left = index * 2 + 1;
        }

    }

    private void heapInsert(int[] nums, int index) {
        while (nums[index] < nums[(index - 1) / 2]) {
            swap(nums, index, (index - 1) / 2);
            index = (index - 1) / 2;
        }
    }

    public ListNode reverse(ListNode root) {
        if (root == null) {
            return root;
        }
        ListNode cur = root;
        ListNode prev = null;
        prev.next = cur;
        while (cur != null) {
            ListNode next = cur.next;
            cur.next = prev;
            prev = cur;
            cur = next;
        }
        return prev.next;
    }

    public ListNode RingNodeInListNode(ListNode root) {
        if (root == null || root.next == null) {
            return null;
        }
        boolean ringFlag = false;
        ListNode fast = root, slow = root;
        while (fast != null && fast.next != null) {
            fast = fast.next.next;
            slow = slow.next;
            if (fast == slow) {
                ringFlag = true;
                ListNode fistMeet = fast;
            }
        }
        if (ringFlag) {
            fast = root;
            while (fast != slow) {
                fast = fast.next;
                slow = slow.next;
            }
            return fast;
        } else {
            return null;
        }
    }


    public String sortByFrequency(String s) {
        if (s == null) {
            return null;
        }
        char[] ch = s.toCharArray();
        HashMap<Character, Integer> map = new HashMap<>();
        for (char c : ch) {
            map.put(c, map.getOrDefault(c, 0) + 1);
        }
        char[] resCh = new char[s.length()];

        int i = 0;
        Set<Character> keySet = map.keySet();
        /*for (Character character : keySet) {
            while (map.get(character)-->=0) {

            resCh[i++] =character ;
            }

        }*/
        return null;
    }
}









