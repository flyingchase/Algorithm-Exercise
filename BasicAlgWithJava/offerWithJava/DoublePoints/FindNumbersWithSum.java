package DoublePoints;

import java.util.ArrayList;
import java.util.Arrays;

/* 
在有序数组中找出两个数，使得和为给定的数 S。如果有多对数字的和等于 S，输出两个数的乘积最小的。


*/
public class FindNumbersWithSum {
    public ArrayList<Integer> FindNumbersWithSum(int[] array, int sum) {
        ArrayList<Integer> reList = new ArrayList<>();
        if (array == null || array.length < 2 || sum <= array[0]) {
            return reList;
        }
        int left = 0, right = array.length - 1;
        while (left < right) {
            int curSum = array[left] + array[right];
            if (curSum == sum) {
                reList.add(array[left]);
                reList.add(array[right]);
                return reList;
            } else if (curSum < sum) {
                left++;
            } else {
                right--;
            }
        }
        Arrays.asList(array[left], array[right]);
        return reList;

    }
}
