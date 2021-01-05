import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @author liamcoder
 * @date 2021/1/5 8:29
 */
public class Solution830 {
    public List<List<Integer>> largeGroupPositions(String s) {
        if (s == null) {
            return null;
        }

        List<List<Integer>> result = new ArrayList<>();
        // 已经连续的字符串长度
        int len = 0;

        int n = s.length();
        for (int i = 0; i < n; ++i) {
            if (i == n - 1 || s.charAt(i) != s.charAt(i + 1)) {
                if (len >= 3) {
                    // abbxxxxzzy   6 - 4 + 1
                    result.add(Arrays.asList(i - len + 1, i));
                }
                len = 1;
            } else {
                ++len;
            }
        }
        return result;
    }
}
