import entity.TreeNode;

/**
 * @author liamcoder
 * @date 2020/12/30 14:30
 */
public class Solution101 {
    public boolean isSymmetric(TreeNode root) {
        return check(root, root);
    }

    private boolean check(TreeNode left, TreeNode right) {
        if (left == null && right == null) {
            return true;
        }
        if (left == null || right == null) {
            return false;
        }
        return left.val == right.val && check(left.left, right.right) && check(right.left, left.right);
    }
}
