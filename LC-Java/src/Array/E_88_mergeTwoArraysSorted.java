package Array;

/*给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

*/
public class E_88_mergeTwoArraysSorted {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int[] res = new int[m + n];
        int p1 = 0, p2 = 0, i = 0;
        while (p1 < m && p2 < n) {
            res[i++] = nums1[p1] <= nums2[p2] ? nums1[p1++] : nums2[p2++];
        }
        while (p1<m) {
            res[i++]=nums1[p1++];
        }
        while (p2<n) {
            res[i++]=nums2[p2++];
        }
        for (int j = 0; j < res.length; j++) {
            nums1[j] = res[j];
        }
    }

}
