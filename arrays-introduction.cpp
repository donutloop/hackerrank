#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    
    int n;
    scanf("%d", &n);
    
    int elements[n];
    
    for (int i = 0; i < n; i++) {
        int element;
        scanf("%d", &element);
        elements[i] = element;  
    }
    
    for (int i = n-1; i >= 0; i--) {
        printf("%d ", elements[i]);  
    }
    
      
    return 0;
}
