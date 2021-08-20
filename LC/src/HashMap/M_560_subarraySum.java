package HashMap;

import java.util.Arrays;
import java.util.HashMap;

/*给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。*/
public class M_560_subarraySum {

    public static void main(String[] args) {
        int[] nums = {4, 5, 0, -2, -3, 1};
        System.out.println(new M_560_subarraySum().subarraySum2(nums, 5));
    }

    public int subarraySum(int[] nums, int k) {

        if (nums == null) {
            return 0;
        }
        HashMap<Integer, Integer> map = new HashMap<>();
        int[] sum = Arrays.copyOf(nums, nums.length);
        for (int i = 1; i < sum.length; i++) {
            sum[i] += sum[i - 1];
        }
        int count = 0;
        map.put(0, 1);
        for (int num : sum) {
            if (map.containsKey(num - k)) {
                count += map.get(num - k);
            }
            map.put(num, map.getOrDefault(num, 0) + 1);
        }

        return count;
    }


    public int subarraySum2(int[] nums, int k) {

        if (nums == null) {
            return 0;
        }
        HashMap<Integer, Integer> map = new HashMap<>();
        int count = 0, sum = 0;
        map.put(0, 1);
        for (int num : nums) {
            sum += num;
            sum %= k;
            if (sum < 0) {
                sum += k;
            }
            count += map.getOrDefault((sum), 0);
            map.put(sum, map.getOrDefault(sum, 0) + 1);
        }
        return count;
    }

}
