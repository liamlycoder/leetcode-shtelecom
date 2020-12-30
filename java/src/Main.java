import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

/**
 * @author luyu
 */
public class Main {
    public static void main(String[] args) {
        int[] nums = {1, 2, 3, 4, 5};
        List<Integer> numberList = Arrays.stream(nums).boxed().collect(Collectors.toList());
    }
}
