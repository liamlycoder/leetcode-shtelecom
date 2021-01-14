import java.util.ArrayList;
import java.util.List;

/**
 * @author liamcoder
 * @date 2021/1/14 8:50
 */
public class Solution1018 {
    public List<Boolean> prefixesDivBy5(int[] A) {
        List<Boolean> result = new ArrayList<>();

        int num = 0;
        for (int i = 0; i < A.length; ++i) {
            num <<= 1;
            num += A[i];
            num %= 10;
            result.add(num % 5 == 0);
        }
        return result;
    }
}
