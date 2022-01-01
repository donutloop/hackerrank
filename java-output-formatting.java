import java.util.Scanner;

public class Solution {

    public static void main(String[] args) {
            Scanner sc=new Scanner(System.in);
            System.out.println("================================");
            for(int i=0;i<3;i++){
                String s1=sc.next();
                int x=sc.nextInt();
                String pad = "";
                int gap = s1.length();
                
                if (x < 10) {
                    pad = "00";
                } else if (x < 100) {
                    pad = "0";
                }
                
                StringBuilder sb = new StringBuilder();
                for (int j = 0; j < 15-gap; j++) {
                    sb.append(' ');
                }
                
                System.out.printf("%s%s%s%d\n", s1, sb, pad, x);
            }
            System.out.println("================================");

    }
}



