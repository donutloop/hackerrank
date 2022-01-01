

    public static String getSmallestAndLargest(String s, int k) {
        String smallest = "";
        String largest = "";
        
        
        
        for (int i = k; i <= s.length(); i++) {
            if (largest.isEmpty() || largest.compareTo(s.substring(i-k, i)) < 0) {
                largest = s.substring(i-k, i);
            }
            if (smallest.isEmpty() || smallest.compareTo(s.substring(i-k, i)) > 0) {
                smallest = s.substring(i-k, i);
            } 
        }
              
        
        return smallest + "\n" + largest;
    }

