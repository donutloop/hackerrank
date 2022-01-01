import java.io.*;
import java.util.*;

public class Solution {

    public static void main(String[] args) {
        
        Scanner sc=new Scanner(System.in);
        String A=sc.next();
        int j = A.length()-1;
        for (int i = 0; i < A.length()/2; i++) {
            if (A.charAt(i) != A.charAt(j)) {
                System.out.println("No");
                return;
            }     
            j--;
        }
        System.out.println("Yes");
    }
}



