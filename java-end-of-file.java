import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int count = 1;
        while(scanner.hasNext()) {
            String s = scanner.nextLine();
            System.out.printf("%d %s \n", count, s);
            count++;
        }            
        scanner.close();
    }
}
