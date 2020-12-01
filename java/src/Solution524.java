import java.util.List;

/**
 * @author liamcoder
 * @date 2020/12/1 14:24
 */
public class Solution524 {
    public String findLongestWord(String s, List<String> d) {
        String result = "";
        for (String target : d) {
            int l1 = result.length();
            int l2 = target.length();
            if (l1 > l2 || (l1 == l2 && result.compareTo(target) < 0)) {
                continue;
            }
            if (isSubStr(s, target)) {
                result = target;
            }
        }
        return result;
    }

    private boolean isSubStr(String s, String target) {
        int i = 0, j = 0;
        while (i < s.length() && j < target.length()) {
            if (s.charAt(i) == target.charAt(j)) {
                j++;
            }
            i++;
        }
        return j == target.length();
    }
}
