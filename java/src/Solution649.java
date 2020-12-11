import java.util.ArrayDeque;
import java.util.Deque;

/**
 * @author liamcoder
 * @date 2020/12/11 15:47
 */
public class Solution649 {
    public String predictPartyVictory(String senate) {
        Deque<Integer> rq = new ArrayDeque<>();
        Deque<Integer> dq = new ArrayDeque<>();
        int n = senate.length();
        for (int i = 0; i < n; ++i) {
            char c = senate.charAt(i);
            if ('R' == c) {
                rq.add(i);
            } else if ('D' == c) {
                dq.add(i);
            }
        }

        while (!rq.isEmpty() && !dq.isEmpty()) {
            int r = rq.remove();
            int d = dq.remove();
            if (r < d) {
                rq.add(r + n);
            } else if (r > d) {
                dq.add(d + n);
            }
        }

        return rq.isEmpty() ? "Dire" : "Radiant";
    }
}
