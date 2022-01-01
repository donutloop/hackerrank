#include <iostream>
#include <cstdio>
using namespace std;

int main() {
    int a, b;
    
    scanf("%d %d", &a, &b);
    
    for (int i = a; i <= b; i++) {
        int k;
        k = 9;
    
        if (i > k && i % 2 == 0) {
            cout << "even \n";
        }else if (i > k && i % 2 == 1) {
            cout << "odd \n";
        }  else if (i == k--) {
            cout << "nine \n";
        } else if (i == k--) {
            cout << "eight \n";
        } else if (i == k--) {
            cout << "seven \n";
        }  else if (i == k--) {
            cout << "six \n";
        }  else if (i == k--) {
            cout << "five \n";
        }  else if (i == k--) {
            cout << "four \n";
        }  else if (i == k--) {
            cout << "three \n";
        }  else if (i == k--) {
            cout << "two \n";
        } else if (i == k--) {
            cout << "one \n";
        }    

    }
    
    return 0;
}
