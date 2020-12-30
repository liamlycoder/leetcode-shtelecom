import entity.TreeNode;

/**
 * @author liamcoder
 * @date 2020/12/30 14:18
 */
public class Solution110 {
    public boolean isBalanced(TreeNode root) {
        if (root == null) {
            return true;
        }
        int leftDepth = getDepth(root.left);
        int rightDepth = getDepth(root.right);
        if (Math.abs(leftDepth - rightDepth) > 1) {
            return false;
        } else {
            return isBalanced(root.left) && isBalanced(root.right);
        }
    }

    private int getDepth(TreeNode root) {
        if (root != null) {
            int leftDepth = getDepth(root.left);
            int rightDepth = getDepth(root.right);
            return Math.max(leftDepth, rightDepth) + 1;
        }
        return 0;
    }
}
