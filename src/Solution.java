public class Solution {
    public static boolean dfs(int[] a, int n, int index, int k, int sum) {
        if (index == n) {
            return sum == k;
        }
        if (dfs(a, n, index + 1, k, sum))
            return true;
        if (dfs(a, n, index + 1, k, sum + a[index]))
            return true;
        return false;
    }

    public static void main(String[] args) {
        int n = 4, k = 13;
        int[] a = new int[] { 1, 2, 4, 7 };
        if (dfs(a, n, 0, k, 0)) {
            System.out.println("Yes");
        } else {
            System.out.println("No");
        }
    }
}