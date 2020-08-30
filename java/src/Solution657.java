import java.util.HashMap;
import java.util.Map;

public class Solution657 {
    public boolean judgeCircle(String moves) {
        Map<Character, Integer> map = new HashMap<>();
        map.put('U', 0);
        map.put('D', 0);
        map.put('L', 0);
        map.put('R', 0);

        for (int i = 0; i < moves.length(); ++i) {
            int new_value = map.get(moves.charAt(i)) + 1;
            map.put(moves.charAt(i), new_value);
        }
        return map.get('U').equals(map.get('D')) && map.get('L').equals(map.get('R'));
    }
}
