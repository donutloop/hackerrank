#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

int main()
{
    int l;
    int k;
	scanf("%d %d", &k, &l);
    
    float n;
    float m;
    scanf("%f %f", &n, &m);
    
    printf("%d %d \n", k+l, k-l);
    
    printf("%.1f %.1f \n", n+m, n-m);
    
    return 0;
}
