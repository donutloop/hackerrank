import java.io.*;
import java.util.*;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Scanner;
import java.math.BigInteger;

public class Solution {

    public static void main(String[] args) {
        
        Scanner scanner = new Scanner(System.in);
        String value = scanner.next();
        scanner.close();
  
        try {
            MessageDigest md = MessageDigest.getInstance("SHA-256");
  
            byte[] messageDigest = md.digest(value.getBytes());
  
            BigInteger no = new BigInteger(1, messageDigest);
  
            String hashtext = no.toString(16);
            while (hashtext.length() < 64) {
                hashtext = "0" + hashtext;
            }
            
            System.out.println(hashtext);
        }    
        catch (NoSuchAlgorithmException e) {
            throw new RuntimeException(e);
        }
    }
}
