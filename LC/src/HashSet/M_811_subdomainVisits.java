package HashSet;

import DataStructure.ListNode;

import javax.imageio.stream.ImageInputStreamImpl;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.HashMap;
import java.util.List;

/*811. 子域名访问计数*/
public class M_811_subdomainVisits {

    public static void main(String[] args) {
        String[] cpdomains = {"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"};
        System.out.println(new M_811_subdomainVisits().subdomainVisits2(cpdomains).toString());
    }
    public List<String> subdomainVisits(String[] cpdomains) {
        List<String> res = new ArrayList<>();

        if (cpdomains==null) {
            return res;
        }
        HashMap<String, Integer> map = new HashMap<String,Integer>();
        for (String s : cpdomains) {
            int i = s.indexOf(' ');
            String num = s.substring(0, i);
            int n = Integer.valueOf(num);
            String domains = s.substring(i + 1);
            for (int j = 0; j < domains.length(); j++) {
                if (domains.charAt(j)=='.') {
                    String domain = domains.substring(j+1);
                    map.put(domain,map.getOrDefault(domain,0)+n);
                }
            }
            map.put(domains,map.getOrDefault(domains,0)+n);
        }

        for (String s : map.keySet()) {
            res.add(map.get(s)+" "+s);
        }
        return res;
    }

    public List<String>subdomainVisits2 (String[] cpdomains) {
        List<String> res = new ArrayList<>();
        if (cpdomains==null) {
            return res;
        }

        HashMap<String, Integer> map = new HashMap<>();
        for (String domains : cpdomains) {
            String[] s = domains.split(" ");
            int n = Integer.valueOf(s[0]);
            String[] domain = s[1].split("\\.");
            String temp = "";
            for (int i = domain.length - 1; i >= 0; i--) {
                temp  = domain[i]+(temp.equals("")?temp:"."+temp);
                map.put(temp,map.getOrDefault(temp,0)+n);
            }
        }

        for (String s : map.keySet()) {
            res.add(map.get(s)+" "+s);
        }
        return res;
    }
}
