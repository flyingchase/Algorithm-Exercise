package String;

import java.util.Arrays;
import java.util.HashSet;
import java.util.List;

public class E_824_toGoatLatin {
    public static void main(String[] args) {
        String s = "I speak Goat Latin";
        System.out.printf("res: ", new E_824_toGoatLatin().toGoatLatin(s));
    }

    private static final HashSet<Character> vowels = new HashSet<>(
            Arrays.asList('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'));

    public String toGoatLatin(String S) {
        if (S == null || S.length() == 0) {
            return "";
        }

        StringBuilder sb = new StringBuilder();
        StringBuilder suffix = new StringBuilder("a");

        for (String w : S.split(" ")) {
            if (sb.length() != 0) {
                sb.append(" ");
            }

            char fChar = w.charAt(0);
            if (vowels.contains(fChar)) {
                sb.append(w);
            } else {
                sb.append(w.substring(1));
                sb.append(fChar);
            }

            sb.append("ma").append(suffix);

            suffix.append("a");
        }
        // System.out.println("sb.toString() = " + sb.toString());
        return sb.toString();
    }
}
