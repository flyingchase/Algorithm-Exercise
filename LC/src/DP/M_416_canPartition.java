package DP;

import java.util.Arrays;

public class M_416_canPartition {
    public boolean canPartition(int[] nums) {
        int sum = Arrays.stream(nums).sum();
        if (sum % 2 != 0) {
            return false;
        }
        sum/=2;
        int n = nums.length;
        boolean[][] dp = new boolean[n + 1][sum + 1];
        for (int i = 0; i <= n; i++) {
            // 目标和 0 不取
            dp[i][0]=true;
        }
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= sum; j++) {
                // 不放入背包 i与 i-1 相同
                dp[i][j]=dp[i-1][j];
                // 放入背包 空间减少
                if (j-nums[i-1]>=0) {
                    dp[i][j]|=dp[i-1][j-nums[j]];
                }
            }
        }
        return dp[n][sum];

    }
}
