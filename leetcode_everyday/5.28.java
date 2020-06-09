package Mengfei.learn;

import java.util.ArrayList;
import java.util.List;

public class LCdecodeString {

    List<String> nums = new ArrayList<String>();

    public String decodeString(String s) {
        if (nums.size() == 0) {
            nums.add("1");
            nums.add("2");
            nums.add("3");
            nums.add("4");
            nums.add("5");
            nums.add("6");
            nums.add("7");
            nums.add("8");
            nums.add("9");
            nums.add("0");
        }
        String[] chars = s.split("");
        String reString = "";
        for (int i = 0; i < chars.length; i++) {
            String ch = chars[i];
            if (nums.contains(ch)) {
                String sNum = ch;
                while (nums.contains(ch) && i < chars.length) {
                    ch = chars[++i];
                    if (nums.contains(ch)) {
                        sNum += ch;
                    }
                }
                String string = kuohao(s.substring(i, s.length()));
                for (int z = 1; z <= Integer.parseInt(sNum); z++) {
                    reString += decodeString(string.substring(1, string.length() - 1));
                }
                i += string.length() - 1;
            } else {
                reString += ch;
            }
        }
        return reString;
    }

    public String kuohao(String string) {
        String[] strings = string.split("");
        int zuo = 1;
        int i = 1;
        for (; i < strings.length; i++) {
            if (strings[i].equals("[")) {
                zuo++;
            } else if (strings[i].equals("]")) {
                zuo--;
            }
            if (zuo == 0) {
                break;
            }
        }
        return string.substring(0, i + 1);
    }
}