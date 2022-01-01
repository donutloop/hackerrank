#include <stdio.h>
#include <stdlib.h>

int main()
{
    int num, *arr, i;
    scanf("%d", &num);
    arr = (int*) malloc(num * sizeof(int));
    for(i = 0; i < num; i++) {
        scanf("%d", arr + i);
    }

    int j;
    j = num - 1;
    for(i = 0; i < num/2; i++) {
        int n;
        n = arr[i];
        int k;
        k = arr[j];
        arr[j] = n;
        arr[i] = k;
        j--;
    }
    
    /* Write the logic to reverse the array. */

    for(i = 0; i < num; i++)
        printf("%d ", *(arr + i));
    return 0;
}
