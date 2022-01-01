import java.util.*;
import java.text.*;
import java.text.NumberFormat;
import java.util.Locale;        

public class Solution {
    
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        double payment = scanner.nextDouble();
        scanner.close();
        
        // Write your code here.
        
        NumberFormat usNf = NumberFormat.getCurrencyInstance(Locale.US);
        String usPayment = usNf.format(payment);
        
        NumberFormat frNf = NumberFormat.getCurrencyInstance(Locale.FRANCE);
        String frPayment = frNf.format(payment);
        
        NumberFormat chinaNf = NumberFormat.getCurrencyInstance(Locale.CHINA);
        String chinaPayment = chinaNf.format(payment);
        
        
        NumberFormat indiaNf = NumberFormat.getCurrencyInstance(new Locale("en", "IN"));
        String indiaPayment = indiaNf.format(payment);
         
        
       System.out.println("US: " +  usPayment);
       System.out.println("India: " + indiaPayment);
       System.out.println("China: " +  chinaPayment);
       System.out.println("France: " +  frPayment);
    }
}
