import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @author liamcoder
 * @date 2020/12/28 15:34
 */
public class Solution347 {
    public int[] topKFrequent(int[] nums, int k) {
        // 保存频率，键为数字，值为该数字对应的频率
        Map<Integer, Integer> frequencies = new HashMap<>();
        for (int num : nums) {
            frequencies.put(num, frequencies.getOrDefault(num, 0) + 1);
        }

        // 设置桶，下标为频率，值为该频率对应的数字一共有哪些
        List<Integer>[] buckets = new ArrayList[nums.length + 1];
        for (Integer key : frequencies.keySet()) {
            int frequency = frequencies.get(key);
            if (buckets[frequency] == null) {
                buckets[frequency] = new ArrayList<>();
            }
            buckets[frequency].add(key);
        }

        // 寻找出现频率前k的数字
        List<Integer> topK = new ArrayList<>();
        for (int i = buckets.length - 1; i >= 0 && k > topK.size(); --i) {
            if (buckets[i] == null) {
                continue;
            }
            if (buckets[i].size() <= k - topK.size()) {
                topK.addAll(buckets[i]);
            } else {
                topK.addAll(buckets[i].subList(0, k - topK.size()));
            }
        }

        // 将结果转换为数组
        int[] res = new int[k];
        for (int i = 0; i < topK.size(); ++i) {
            res[i] = topK.get(i);
        }

        return res;
    }
}
