public class Solution214 {
    public String shortestPalindrome(String s) {
        String result = "";
        int len = s.length();
        if (len == 0) {
            return result;
        }
        String rs = reverse(s);

        int i = 0;
        while (true) {
            if (rs.substring(i, s.length() - 1).equals(s.substring(0, s.length()-1-i))) {
                return rs.substring(0, i) + s;
            }
            i++;
        }
    }

    public String reverse(String s) {
        StringBuilder result = new StringBuilder("");

        for (int i = s.length() - 1; i >= 0; --i) {
            result.append(s.charAt(i));
        }

        return result.toString();
    }
}
