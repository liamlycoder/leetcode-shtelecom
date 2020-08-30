public class Solution557 {
    public String reverseWords(String s) {
        String[] strings = s.split(" ");
        StringBuilder result = new StringBuilder();
        for (int i = 0; i < strings.length; ++i) {
            strings[i] = new StringBuffer(strings[i]).reverse().toString();
            result.append(strings[i]);
            if (i != strings.length - 1) {
                result.append(" ");
            }
        }
        return result.toString();
    }
}
