import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

/* 
给定两个数组，编写一个函数来计算它们的交集。

 

示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
 

说明：

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。
 */
public class UmInTwoList {
    public static int[] intersection1(int[] nums1, int[] nums2) {
        Set<Integer> set1 =new HashSet<>();
        Set<Integer> set2 =new HashSet<>();
        for(int i =0;i<nums1.length;i++) {
            set1.add(nums1[i]);
        }
        for(int i=0;i<nums2.length;i++) {
            if(set1.contains(nums2[i])) {
                set2.add(nums2[i]);
            }
        }
        int[] arr = new int[set2.size()];
        Object[] arr1 = set2.toArray();
        for(int i =0;i<arr1.length;i++) {
            arr[i] =(int)arr1[i];
        }
        return arr;
    }
    
    public static int[] intersection2(int[] nums1, int[] nums2) {
        Arrays.sort(nums1);
        Arrays.sort(nums2);
        int l1=nums1.length,l2=nums2.length;
        int [] helper = new int[l1+l2];
        int index=0,index1=0,index2=0;
        while (index1<l1&&index2<l2) {
            if(nums1[index1]==nums2[index2]) {
                if(index==0||nums1[index1]!=helper[index-1]) {
                    helper[index++] = nums1[index1];
                }
                index1++;
                index2++;
            }
            else if(nums1[index1]<nums2[index2]) {
                index1++;
            }
            else {
                index2++;
            }
        }

        return Arrays.copyOfRange(helper, 0, index);
    }

    public static void main(String[] args) {
        int nums1[] ={1,2,2,1};
        int nums2[] ={2,2};
        System.out.println(Arrays.toString(intersection2(nums1, nums2)));

        
    }
}
