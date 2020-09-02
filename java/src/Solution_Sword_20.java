public class Solution_Sword_20 {
    public boolean isNumber(String s) {
        if (s == null || s.length() == 0)  return false;
        s = s.trim();  // 去掉空格
        boolean numFlag = false;
        boolean dotFlag = false;
        boolean eFlag = false;

        for (int i = 0; i < s.length(); ++i) {
            if (s.charAt(i) >= '0' && s.charAt(i) <= '9') {
                numFlag = true;
            } else if (s.charAt(i) == '.' && !dotFlag && !eFlag) {
                dotFlag = true;
            } else if ((s.charAt(i) == 'e' || s.charAt(i) == 'E') && !eFlag && numFlag) {
                eFlag = true;
                numFlag = false;
            } else if ((s.charAt(i) == '+' || s.charAt(i) == '-') && (i == 0 || s.charAt(i-1) == 'e' || s.charAt(i-1) == 'E')) {

            } else {
                return false;
            }
        }
        return numFlag;
    }
}
