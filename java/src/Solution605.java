/**
 * @author liamcoder
 * @date 2021/1/1 11:13 下午
 */
public class Solution605 {
    public boolean canPlaceFlowers(int[] flowerbed, int n) {
        // 由于条件限制，所以本身就不可能连续两个1，所以可以每隔两个跳一次
        for (int i = 0; i < flowerbed.length; i += 2) {
            if (flowerbed[i] == 0) {
                if (i == flowerbed.length - 1 || flowerbed[i+1] == 0) {
                    --n;
                } else {
                    ++i;
                }
            }
        }
        return n <= 0;
    }
}
