#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    
    int n;
    scanf("%d", &n);
    
    vector<int>v;
    for (int i = 0; i < n; i++) {
        int element;
        scanf("%d", &element);
        v.push_back(element);
    }
    sort(v.begin(),v.end()); 
    
    for(auto value: v) { 
       printf("%d ", value);  
    }
    
    return 0;
}
