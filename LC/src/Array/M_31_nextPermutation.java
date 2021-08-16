package Array;

/*实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列（即，组合出下一个更大的整数）*/
public class M_31_nextPermutation {
    public void nextPermutation(int[] A) {
        if (A == null || A.length <= 1) {
            return;
        }
        int i = A.length - 2;
        while (i >= 0 && A[i] >= A[i + 1]) {
            i--; // Find 1st id i that breaks descending order
        }
        if (i >= 0) {                           // If not entirely descending
            int j = A.length - 1;              // Start from the end
            while (A[j] <= A[i]) {
                j--;           // Find rightmost first larger id j
            }
            swap(A, i, j);                     // Switch i and j
        }
        reverse(A, i + 1, A.length - 1);       // Reverse the descending sequence
    }

    public void swap(int[] A, int i, int j) {
        int tmp = A[i];
        A[i] = A[j];
        A[j] = tmp;
    }

    public void reverse(int[] A, int i, int j) {
        while (i < j) {
            swap(A, i++, j--);
        }
    }
}
