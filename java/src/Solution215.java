/**
 * @author liamcoder
 * @date 2020/12/2 9:40
 */
public class Solution215 {
    public int findKthLargest(int[] nums, int k) {
        k = nums.length - k;
        int l = 0, r = nums.length - 1;
        while (l < r) {
            int j = partition(nums, l, r);
            if (j == k) {
                break;
            } else if (j > k) {
                r = j - 1;
            } else {
                l = j + 1;
            }
        }
        return nums[k];
    }
    private int partition(int[] nums, int l, int r) {
        int i = l, j = r + 1;
        while (true) {
            while (nums[++i] < nums[l] && i < r) ;
            while (nums[--j] > nums[l] && j > l) ;
            if (i >= j) {
                break;
            }
            swap(nums, i, j);
        }
        swap(nums, l, j);
        return j;
    }
    private void swap(int[] a, int i, int j) {
        int t = a[i];
        a[i] = a[j];
        a[j] = t;
    }
}
