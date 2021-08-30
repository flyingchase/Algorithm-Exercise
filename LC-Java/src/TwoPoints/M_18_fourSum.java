package TwoPoints;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class M_18_fourSum {
    public static void main(String[] args) {
        int[] nums = {1,0,-1,0,-2,2};
        System.out.println(new M_18_fourSum().fourSum(nums, 0));
    }
    public List<List<Integer>> fourSum(int[] nums, int target) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(nums);
        if (nums == null || nums.length < 4) {
            return res;
        }

        for (int i = 0; i < nums.length - 3; i++) {
            if (i>0&&nums[i]==nums[i-1]){
                continue;
            }
            for (int j = i + 1; j < nums.length - 2; j++) {
                if (j>i+1&&nums[j]==nums[j-1]) {
                    continue;
                }
                int l = j + 1, r = nums.length - 1;

                while (l < r) {
                    int sum = nums[i] + nums[j] + nums[l] + nums[r];
                    if (sum == target) {
                        res.add(Arrays.asList(nums[i], nums[j], nums[l], nums[r]));
                        while (l < r && nums[l] == nums[l + 1]) {
                            l++;
                        }
                        while (l < r && nums[r] == nums[r - 1]) {
                            r--;
                        }
                        l++;
                        r--;
                    } else if (sum>target) {
                        r--;
                    }else {
                        l++;
                    }
                }
            }
        }
        return res;
    }
}
