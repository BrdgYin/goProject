import java.util.HashMap;
import java.util.Map;

public class Main {
    public static int solve2(int[] work, int day) {
        int[] dp = new int[day + 1]; // 表示第i天最大的作业量
        dp[0] = 0; // 第0天任务量为0
        dp[1] = Math.max(work[0], work[1]); // 第一天的任务量

        for (int i = 2; i <= day; i++) {
            if (i - 2 > 0) { // 来自于休息一天1 [2] 3
                dp[i] = Math.max(dp[i], dp[i - 2] + work[i - 1]);
            }
            if (i - 3 > 0) { // 来自于休息两天1 [2, 3] 4
                dp[i] = Math.max(dp[i], dp[i - 3] + work[i - 1]);
            }
        }
        return dp[day];
    }

    public static int solve(String S, int n) {
        StringBuffer sb = new StringBuffer();
        Map<Character, Integer> map = new HashMap<>();
        int size = S.length();
        int max = Integer.MIN_VALUE;
        int min = Integer.MAX_VALUE;
        for (int i = 0; i < n / size + 1; i++) {
            sb.append(S);
        }
        String res = sb.toString().substring(0, n + 1);

        // 交给map管理
        for (int i = 0; i < n; i++) {
            if (!map.containsKey(res.charAt(i))) {
                map.put(res.charAt(i), 1);
            } else {
                int temp = map.get(res.charAt(i));
                temp = temp + 1;
                map.put(res.charAt(i), temp);
            }
        }

        for (Character key : map.keySet()) {
            if (map.get(key) > max) {
                max = map.get(key);
            }
            if (map.get(key) < min) {
                min = map.get(key);
            }
        }

        return max - min;
    }

    public static void main(String[] args) {
        int[] work = new int[] { 2, 1, 2, 1, 2 };
        int day = 5;
        System.out.println(solve2(work, day));
    }
}