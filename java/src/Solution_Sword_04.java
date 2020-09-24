public class Solution_Sword_04 {
    public boolean findNumberIn2DArray(int[][] matrix, int target) {
        if (matrix == null || matrix.length == 0) {
            return false;
        }
        int m = matrix.length;
        int n = matrix[0].length;
        int row = 0, col = n-1;
        while (row < m && col >= 0) {
            if (target > matrix[row][col]) {
                ++row;
            } else if (target < matrix[row][col]) {
                --col;
            } else {
                return true;
            }
        }

        return false;
    }
}
