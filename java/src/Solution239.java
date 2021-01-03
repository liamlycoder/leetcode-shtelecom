import java.util.Comparator;
import java.util.Deque;
import java.util.LinkedList;
import java.util.PriorityQueue;

/**
 * @author liamcoder
 * @date 2021/1/2 11:06 上午
 */
public class Solution239 {
    public int[] maxSlidingWindow(int[] nums, int k) {
        int len = nums.length;
        PriorityQueue<int[]> pq = new PriorityQueue<>(new Comparator<int[]>() {
            @Override
            public int compare(int[] o1, int[] o2) {
                return o1[0] != o2[0] ? o2[0] - o1[0] : o2[1] - o2[1];
            }
        });

        for (int i = 0; i < k; ++i) {
            pq.offer(new int[]{nums[i], i});
        }

        int[] result = new int[len - k + 1];
        result[0] = pq.peek()[0];
        for (int i = k; i < len; ++i) {
            pq.offer(new int[]{nums[i], i});
            while (pq.peek()[1] <= i - k) {
                pq.poll();
            }
            result[i - k + 1] = pq.peek()[0];
        }
        return result;
    }

    public int[] maxSlidingWindow2(int[] nums, int k) {
        int len = nums.length;
        Deque<Integer> deque = new LinkedList<>();
        int[] result = new int[len - k + 1];

        for (int i = 0; i < len; ++i) {
            while (!deque.isEmpty() && nums[i] >= nums[deque.peekLast()]) {
                deque.pollLast();
            }
            deque.offerLast(i);
            while (deque.peekFirst() <= i - k) {
                deque.pollFirst();
            }
            if (i + 1 >= k) {
                result[i - k + 1] = nums[deque.peekFirst()];
            }
        }
        return result;
    }
}
