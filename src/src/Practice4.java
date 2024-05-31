import java.util.*;


public class Practice4 {

    public static List<Integer> maxSubarray(List<Integer> arr) {
        int maxSubArray = Integer.MIN_VALUE;
        int maxSubsequence = 0;
        int currentSum = 0;
        int maxValue = Integer.MIN_VALUE;
        for (int i : arr) {
            if (i > 0) maxSubsequence += i;
            currentSum += i;
            maxSubArray = Math.max(maxSubArray, currentSum);
            currentSum = Math.max(currentSum, 0);
            maxValue = Math.max(maxValue, i);
        }
        if (maxSubsequence == 0) maxSubsequence = maxValue;
        return Arrays.asList(maxSubArray, maxSubsequence);
    }


    public static int cost(List<Integer> B) {
        int n = B.size();
        int[] dpLow = new int[n];
        int[] dpHigh = new int[n];

        dpLow[0] = 0;
        dpHigh[0] = 0;

        for (int i = 1; i < n; i++) {
            dpLow[i] = Math.max(dpLow[i - 1], dpHigh[i - 1] + Math.abs(1 - B.get(i - 1)));
            dpHigh[i] = Math.max(dpLow[i - 1] + Math.abs(B.get(i) - 1), dpHigh[i - 1] + Math.abs(B.get(i) - B.get(i - 1)));
        }
        return Math.max(dpLow[n - 1], dpHigh[n - 1]);

    }

    public static void main(String[] args) {

    }


    public class TreeNode {
        int val;
        Leetcode3.TreeNode left;
        Leetcode3.TreeNode right;

        TreeNode() {
        }

        TreeNode(int val) {
            this.val = val;
        }

        TreeNode(int val, Leetcode3.TreeNode left, Leetcode3.TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }


    static class Node {
        public int val;
        public Leetcode3.Node left;
        public Leetcode3.Node right;
        public Leetcode3.Node next;

        public Node() {
        }

        public Node(int _val) {
            val = _val;
        }

        public Node(int _val, Leetcode3.Node _left, Leetcode3.Node _right, Leetcode3.Node _next) {
            val = _val;
            left = _left;
            right = _right;
            next = _next;
        }
    }


}
