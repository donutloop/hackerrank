#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

int main() {
	
    int n;
    scanf("%d", &n);
    
    int sum;
    for(;;) {
    
        int r;
        r = n % 10;
        n = n / 10;
        sum += r;
        
        if (n < 9) {
            sum += n;
            break;   
        }
    }
    
    printf("%d", sum);
    
    return 0;
}
