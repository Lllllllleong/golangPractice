import java.util.ArrayDeque;
import java.util.Deque;

public class Leetcode2 {




    public int[] maxSlidingWindow(int[] nums, int k) {
        if (k == 1) return nums;
        int n = nums.length;
        int[] output = new int[n-k+1];
        int outputIndex = 0;
        Deque<Integer> windowQueue = new ArrayDeque<>();
        windowQueue.addLast(Integer.MIN_VALUE);
        for (int i = 0; i < nums.length; i++) {
            int current = nums[i];
            int windowLB = i - k + 1;
            while (!windowQueue.isEmpty() && windowQueue.peek() < windowLB) windowQueue.pop();

            //New max
            if (windowQueue.isEmpty() || current >= nums[windowQueue.peek()]) {
                windowQueue.clear();
                windowQueue.add(i);
            } else {
                while (nums[windowQueue.peekLast()] < current) {
                    windowQueue.pollLast();
                }
                windowQueue.addLast(i);
            }
            //Begin adding to output
            if (i >= (k-1)) {
                output[outputIndex] = nums[windowQueue.peek()];
                outputIndex++;
            }
        }
        return output;
    }
}
