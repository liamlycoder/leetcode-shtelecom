import java.util.HashMap;
import java.util.Map;

/**
 * @author liamcoder
 * @date 2020/12/16 8:45
 */
public class Solution290 {
    public boolean wordPattern(String pattern, String s) {
        Map<String, Character> str2ch = new HashMap<>();
        Map<Character, String> ch2str = new HashMap<>();
        int i = 0, m = s.length();
        for (int p = 0; p < pattern.length(); ++p) {
            char ch = pattern.charAt(p);
            if (i >= m) {
                return false;
            }
            int j = i;
            while (j < m && s.charAt(j) != ' ') {
                ++j;
            }
            String tmp = s.substring(i, j);
            if (str2ch.containsKey(tmp) && str2ch.get(tmp) != ch || ch2str.containsKey(ch) && !tmp.equals(ch2str.get(ch))) {
                return false;
            }
            str2ch.put(tmp, ch);
            ch2str.put(ch, tmp);
            i = j + 1;
        }
        return i >= m;
    }
}
