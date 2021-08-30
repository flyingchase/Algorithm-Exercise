package TwoPoints;

import com.sun.source.tree.Tree;

import javax.swing.text.html.HTML;
import java.sql.Array;
import java.util.*;

/*给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。*/
public class M_16_threeSumClosest {
    public static void main(String[] args) {
        int[] nums = {2,1,-3,0};
        System.out.println(new M_16_threeSumClosest().threeSumClosest(nums, 1));
    }

    public int threeSumClosest(int[] nums, int target) {
        ArrayList<Integer> res = new ArrayList<>();
        Arrays.sort(nums);
        if (nums.length<3) {
            return 0;
        }
        for (int i = 0; i < nums.length-2; i++) {
            int l=i+1,r=nums.length-1;
            while (l<r) {
                int sum = nums[i]+nums[l]+nums[r];
                res.add(sum);
                if (sum>target) {

                r--;
                }else if (sum< target){

                l++;
                }
            }
        }
        int mingap =Integer.MAX_VALUE;

        HashMap<Integer, Integer> treeMap = new HashMap<>();
        for (Integer num : res) {
            treeMap.put(Math.abs(num-target),num);
            mingap=Math.min(Math.abs(num-target),mingap);
        }
        return treeMap.get(mingap);
    }
}
