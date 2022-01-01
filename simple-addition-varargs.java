

class Add  {
    void add(int... values) {
        int sum = 0;
        int count = 0;
        String s = "";
        for (int value : values) {
            sum += value;
            if (count == 0) {
              s += value;        
            } else {
              s += "+" + value;  
            }
            count++;
        }
        s += "=" + sum;
        System.out.println(s);
        return;
    }
}



