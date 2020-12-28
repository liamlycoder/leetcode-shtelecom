import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @author liamcoder
 * @date 2020/12/28 16:34
 */
public class Solution451 {
    public String frequencySort(String s) {
        // 保存每个字符出现的频率，键为字符，值为对应的频率
        Map<Character, Integer> map = new HashMap<>();
        for (char c : s.toCharArray()) {
            map.put(c, map.getOrDefault(c, 0) + 1);
        }

        // 设置桶，下标为频率，值为该频率的所有字符
        List<Character>[] buckets = new ArrayList[s.length() + 1];
        for (Character key : map.keySet()) {
            int f = map.get(key);
            if (buckets[f] == null) {
                buckets[f] = new ArrayList<>();
            }
            buckets[f].add(key);
        }

        // 从桶里倒序提取出字符，即可完成从大到小排序
        StringBuilder stringBuilder = new StringBuilder();
        for (int i = buckets.length - 1; i >= 0; --i) {
            if (buckets[i] == null) {
                continue;
            }
            for (Character c : buckets[i]) {
                for (int j = 0; j < i; ++j) {
                    stringBuilder.append(c);
                }
            }
        }
        return stringBuilder.toString();
    }
}
